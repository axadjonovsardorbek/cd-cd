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

type DepartmentRepo struct {
	db *sql.DB
}

func NewDepartmentRepo(db *sql.DB) *DepartmentRepo {
	return &DepartmentRepo{db: db}
}

func (d *DepartmentRepo) Create(req *cp.DepartmentCreateReq) (*cp.DepartmentRes, error) {
	id := uuid.New().String()
	department := cp.DepartmentRes{}

	query := `
	INSERT INTO departments(
		id,
		name,
		company_id
	) VALUES ($1, $2, $3)
	RETURNING
		id,
		name,
		company_id
	`

	row := d.db.QueryRow(query, id, req.Name, req.CompanyId)

	err := row.Scan(
		&department.Id,
		&department.Name,
		&department.CompanyId,
	)

	if err != nil {
		log.Println("Error while creating department: ", err)
		return nil, err
	}

	log.Println("Successfully created department")

	return &department, nil
}

func (d *DepartmentRepo) GetById(id *cp.Byid) (*cp.DepartmentGetByIdRes, error) {
	department := cp.DepartmentGetByIdRes{
		Department: &cp.DepartmentRes{},
	}

	query := `
	SELECT 
		id,
		name,
		company_id
	FROM 
		departments
	WHERE 
		id = $1
	AND 
		deleted_at = 0	
	`

	row := d.db.QueryRow(query, id.Id)

	err := row.Scan(
		&department.Department.Id,
		&department.Department.Name,
		&department.Department.CompanyId,
	)


	if err != nil {
		log.Println("Error while getting department by id: ", err)
		return nil, err
	}

	log.Println("Successfully got department")

	return &department, nil
}

func (d *DepartmentRepo) GetAll(req *cp.DepartmentGetAllReq) (*cp.DepartmentGetAllRes, error) {
	departments := cp.DepartmentGetAllRes{}

	query := `
	SELECT 
		id,
		name,
		company_id
	FROM 
		departments
	WHERE 
		deleted_at = 0
	`

	var conditions []string
	var args []interface{}

	if req.Name != "" && req.Name != "string" {
		conditions = append(conditions, " LOWER(name) LIKE LOWER($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, req.Name)
	}
	if req.CompanyId != "" && req.CompanyId != "string" {
		conditions = append(conditions, " company_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.CompanyId)
	}
	if len(conditions) != 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}


	var defaultLimit int64
	err := d.db.QueryRow("SELECT COUNT(1) FROM departments WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		log.Println("Error while getting count : ", err)
		return nil, err
	}
	if req.Filter.Limit == 0 {
		req.Filter.Limit = defaultLimit
	}

	args = append(args, req.Filter.Limit, req.Filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := d.db.Query(query, args...)

	if err != nil {
		log.Println("Error while retriving departments: ", err)
		return nil, err
	}

	for rows.Next() {
		department := cp.DepartmentRes{}

		err := rows.Scan(
			&department.Id,
			&department.Name,
			&department.CompanyId,
		)

		if err != nil {
			log.Println("Error while scanning all departments: ", err)
			return nil, err
		}

		departments.Departments = append(departments.Departments, &department)
	}

	log.Println("Successfully fetched all departments")

	return &departments, nil
}

func (d *DepartmentRepo) Update(req *cp.DepartmentUpdateReq) (*cp.DepartmentRes, error) {
	department := cp.DepartmentRes{}

	query := `
	UPDATE
		departments
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
	if req.CompanyId != "" && req.CompanyId != "string" {
		conditions = append(conditions, " company_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.CompanyId)
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
		company_id
	`

	args = append(args, req.Id)

	row := d.db.QueryRow(query, args...)

	err := row.Scan(
		&department.Id,
		&department.Name,
		&department.CompanyId,
	)

	if err != nil {
		log.Println("Error while updating department: ", err)
		return nil, err
	}

	log.Println("Successfully updated department")

	return &department, nil
}

func (d *DepartmentRepo) Delete(id *cp.Byid) (*cp.Void, error) {
	query := `
	UPDATE 
		departments
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := d.db.Exec(query, id.Id)

	if err != nil {
		log.Println("Error while deleting department: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("department with id %s not found", id.Id)
	}

	log.Println("Successfully deleted department")

	return &cp.Void{
		Success: true,
		Message: "Successfully deleted department",
	}, nil
}
