package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	cp "company/genprotos"

	"github.com/google/uuid"
)

type CompanyRepo struct {
	db *sql.DB
}

func NewCompanyRepo(db *sql.DB) *CompanyRepo {
	return &CompanyRepo{db: db}
}

func (c *CompanyRepo) Create(req *cp.CompanyCreateReq) (*cp.CompanyRes, error) {

	id := uuid.New().String()
	company := cp.CompanyRes{}

	query := `
	INSERT INTO companies(
		id,
		name,
		address,
		industry,
		website
	) VALUES ($1, $2, $3, $4, $5)
	RETURNING
		id,
		name,
		address,
		industry,
		website
	`

	row := c.db.QueryRow(query, id, req.Name, req.Address, req.Industry, req.Website)

	err := row.Scan(
		&company.Id,
		&company.Name,
		&company.Address,
		&company.Industry,
		&company.Website,
	)

	if err != nil {
		log.Println("Error while creating company: ", err)
		return nil, err
	}

	log.Println("Successfully created company")

	return &company, nil
}

func (c *CompanyRepo) GetById(id *cp.Byid) (*cp.CompanyGetByIdRes, error) {
	company := cp.CompanyGetByIdRes{
		Company: &cp.CompanyRes{},
	}

	query := `
	SELECT 
		id,
		name,
		address,
		industry,
		website
	FROM 
		companies
	WHERE 
		id = $1
	AND 
		deleted_at = 0	
	`

	row := c.db.QueryRow(query, id.Id)

	err := row.Scan(
		&company.Company.Id,
		&company.Company.Name,
		&company.Company.Address,
		&company.Company.Industry,
		&company.Company.Website,
	)


	if err != nil {
		log.Println("Error while getting company by id: ", err)
		return nil, err
	}

	log.Println("Successfully got company")

	return &company, nil
}

func (c *CompanyRepo) GetAll(req *cp.CompanyGetAllReq) (*cp.CompanyGetAllRes, error) {
	companies := cp.CompanyGetAllRes{}

	query := `
	SELECT 
		id,
		name,
		address,
		industry,
		website
	FROM 
		companies
	WHERE 
		deleted_at = 0
	`

	var conditions []string
	var args []interface{}

	if req.Name != "" && req.Name != "string" {
		conditions = append(conditions, " LOWER(name) LIKE LOWER($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, "%" + req.Name + "%")
	}
	if req.Address != "" && req.Address != "string" {
		conditions = append(conditions, " LOWER(address) LIKE LOWER($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, "%" + req.Address + "%")
	}
	if req.Industry != "" && req.Industry != "string" {
		conditions = append(conditions, " LOWER(industry) LIKE LOWER($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, "%" + req.Industry + "%")
	}
	if req.Website != "" && req.Website != "string" {
		conditions = append(conditions, " website = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Website)
	}
	if len(conditions) != 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	var defaultLimit int64
	err := c.db.QueryRow("SELECT COUNT(1) FROM companies WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		log.Println("Error while getting count : ", err)
		return nil, err
	}
	if req.Filter.Limit == 0 {
		req.Filter.Limit = defaultLimit
	}

	args = append(args, req.Filter.Limit, req.Filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := c.db.Query(query, args...)

	if err != nil {
		log.Println("Error while retriving companies: ", err)
		return nil, err
	}

	for rows.Next() {
		company := cp.CompanyRes{}

		err := rows.Scan(
			&company.Id,
			&company.Name,
			&company.Address,
			&company.Industry,
			&company.Website,
		)

		if err != nil {
			log.Println("Error while scanning all companies: ", err)
			return nil, err
		}

		companies.Companies = append(companies.Companies, &company)
	}

	log.Println("Successfully fetched all companies")

	return &companies, nil
}

func (c *CompanyRepo) Update(req *cp.CompanyUpdateReq) (*cp.CompanyRes, error) {
	company := cp.CompanyRes{}

	query := `
	UPDATE
		companies
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.Id == "" || req.Id == "string" {
		return nil, errors.New("nothing to update")
	}
	if req.Name != "" && req.Name != "string" {
		conditions = append(conditions, " name = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Name)
	}
	if req.Address != "" && req.Address != "string" {
		conditions = append(conditions, " address = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Address)
	}
	if req.Industry != "" && req.Industry != "string" {
		conditions = append(conditions, " industry = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Industry)
	}
	if req.Website != "" && req.Website != "string" {
		conditions = append(conditions, " website = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Website)
	}
	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0 "

	query += `
	RETURNING
		id,
		name,
		address,
		industry,
		website
	`

	args = append(args, req.Id)

	row := c.db.QueryRow(query, args...)

	err := row.Scan(
		&company.Id,
		&company.Name,
		&company.Address,
		&company.Industry,
		&company.Website,
	)

	if err != nil {
		log.Println("Error while updating company: ", err)
		return nil, err
	}

	log.Println("Successfully updated company")

	return &company, nil
}

func (c *CompanyRepo) Delete(id *cp.Byid) (*cp.Void, error) {

	query := `
	UPDATE 
		companies
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := c.db.Exec(query, id.Id)

	if err != nil {
		log.Println("Error while deleting company: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("company with id %s not found", id.Id)
	}

	log.Println("Successfully deleted company")

	return &cp.Void{
		Success: true,
		Message: "Successfully deleted company",
	}, nil
}
