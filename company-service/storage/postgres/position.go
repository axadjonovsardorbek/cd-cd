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

type PositionRepo struct {
	db *sql.DB
}

func NewPositionRepo(db *sql.DB) *PositionRepo {
	return &PositionRepo{db: db}
}

func (p *PositionRepo) Create(req *cp.PositionCreateReq) (*cp.PositionRes, error) {
	id := uuid.New().String()
	position := cp.PositionRes{}

	query := `
	INSERT INTO positions(
		id,
		title,
		description,
		department_id
	) VALUES ($1, $2, $3, $4)
	RETURNING
		id,
		title,
		description,
		department_id
	`

	row := p.db.QueryRow(query, id, req.Title, req.Description, req.DepartmentId)

	err := row.Scan(
		&position.Id,
		&position.Title,
		&position.Description,
		&position.DepartmentId,
	)

	if err != nil {
		log.Println("Error while creating position: ", err)
		return nil, err
	}

	log.Println("Successfully created position")

	return &position, nil
}

func (p *PositionRepo) GetById(id *cp.Byid) (*cp.PositionGetByIdRes, error) {
	position := cp.PositionGetByIdRes{
		Position: &cp.PositionRes{},
	}

	query := `
	SELECT 
		id,
		title,
		description,
		department_id
	FROM 
		positions
	WHERE 
		id = $1
	AND 
		deleted_at = 0	
	`

	row := p.db.QueryRow(query, id.Id)

	err := row.Scan(
		&position.Position.Id,
		&position.Position.Title,
		&position.Position.Description,
		&position.Position.DepartmentId,
	)


	if err != nil {
		log.Println("Error while getting position by id: ", err)
		return nil, err
	}

	log.Println("Successfully got position")

	return &position, nil
}

func (p *PositionRepo) GetAll(req *cp.PositionGetAllReq) (*cp.PositionGetAllRes, error) {
	positions := cp.PositionGetAllRes{}

	query := `
	SELECT 
		id,
		title,
		description,
		department_id
	FROM 
		positions
	WHERE 
		deleted_at = 0
	`

	var conditions []string
	var args []interface{}

	if req.DepartmentId != "" && req.DepartmentId != "string" {
		conditions = append(conditions, " department_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.DepartmentId)
	}
	if len(conditions) != 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}


	var defaultLimit int64
	err := p.db.QueryRow("SELECT COUNT(1) FROM positions WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		log.Println("Error while getting count : ", err)
		return nil, err
	}
	if req.Filter.Limit == 0 {
		req.Filter.Limit = defaultLimit
	}

	args = append(args, req.Filter.Limit, req.Filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := p.db.Query(query, args...)

	if err != nil {
		log.Println("Error while retriving positions: ", err)
		return nil, err
	}

	for rows.Next() {
		position := cp.PositionRes{}

		err := rows.Scan(
			&position.Id,
			&position.Title,
			&position.Description,
			&position.DepartmentId,
		)

		if err != nil {
			log.Println("Error while scanning all positions: ", err)
			return nil, err
		}

		positions.Positions = append(positions.Positions, &position)
	}

	log.Println("Successfully fetched all positions")

	return &positions, nil
}

func (p *PositionRepo) Update(req *cp.PositionUpdateReq) (*cp.PositionRes, error) {
	position := cp.PositionRes{}

	query := `
	UPDATE
		positions
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.Id == "" || req.Id == "string" {
		return nil, errors.New("nothing to update")
	}
	if req.Title != "" && req.Title != "string" {
		conditions = append(conditions, " title = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Title)
	}
	if req.Description != "" && req.Description != "string" {
		conditions = append(conditions, " description = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Description)
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
		title,
		description,
		department_id
	`

	args = append(args, req.Id)

	row := p.db.QueryRow(query, args...)

	err := row.Scan(
		&position.Id,
		&position.Title,
		&position.Description,
		&position.DepartmentId,
	)

	if err != nil {
		log.Println("Error while updating position: ", err)
		return nil, err
	}

	log.Println("Successfully updated position")

	return &position, nil
}

func (p *PositionRepo) Delete(id *cp.Byid) (*cp.Void, error) {
	query := `
	UPDATE 
		positions
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := p.db.Exec(query, id.Id)

	if err != nil {
		log.Println("Error while deleting position: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("position with id %s not found", id.Id)
	}

	log.Println("Successfully deleted position")

	return &cp.Void{
		Success: true,
		Message: "Successfully deleted position",
	}, nil
}
