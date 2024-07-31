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

type ResumeRepo struct {
	db *sql.DB
}

func NewResumeRepo(db *sql.DB) *ResumeRepo {
	return &ResumeRepo{db: db}
}

func (r *ResumeRepo) Create(req *cp.ResumeCreateReq) (*cp.ResumeRes, error) {
	id := uuid.New().String()
	resume := cp.ResumeRes{}

	query := `
	INSERT INTO resumes(
		id,
		user_id,
		education,
		summary
	) VALUES ($1, $2, $3, $4)
	RETURNING
		id,
		user_id,
		education,
		summary
	`

	row := r.db.QueryRow(query, id, req.UserId, req.Education, req.Summary)

	err := row.Scan(
		&resume.Id,
		&resume.UserId,
		&resume.Education,
		&resume.Summary,
	)

	if err != nil {
		log.Println("Error while creating resume: ", err)
		return nil, err
	}

	log.Println("Successfully created resume")

	return &resume, nil
}

func (r *ResumeRepo) GetById(id *cp.Byid) (*cp.ResumeGetByIdRes, error) {
	resume := cp.ResumeGetByIdRes{
		Resume: &cp.ResumeRes{},
	}

	query := `
	SELECT 
		id,
		user_id,
		education,
		summary
	FROM 
		resumes
	WHERE 
		id = $1
	AND 
		deleted_at = 0	
	`

	row := r.db.QueryRow(query, id.Id)

	err := row.Scan(
		&resume.Resume.Id,
		&resume.Resume.UserId,
		&resume.Resume.Education,
		&resume.Resume.Summary,
	)

	if err != nil {
		log.Println("Error while getting resume by id: ", err)
		return nil, err
	}

	log.Println("Successfully got resume")

	return &resume, nil
}

func (r *ResumeRepo) GetAll(req *cp.ResumeGetAllReq) (*cp.ResumeGetAllRes, error) {
	resumes := cp.ResumeGetAllRes{}

	query := `
	SELECT 
		id,
		user_id,
		education,
		summary
	FROM 
		resumes
	WHERE 
		deleted_at = 0
	`

	var conditions []string
	var args []interface{}

	if req.UserId != "" && req.UserId != "string" {
		conditions = append(conditions, " user_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.UserId)
	}
	if req.Education != "" && req.Education != "string" {
		conditions = append(conditions, " LOWER(education) LIKE LOWER($"+strconv.Itoa(len(args)+1)+")")
		args = append(args, req.Education)
	}
	if len(conditions) != 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}


	var defaultLimit int64
	err := r.db.QueryRow("SELECT COUNT(1) FROM resumes WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		log.Println("Error while getting count : ", err)
		return nil, err
	}
	if req.Filter.Limit == 0 {
		req.Filter.Limit = defaultLimit
	}

	args = append(args, req.Filter.Limit, req.Filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := r.db.Query(query, args...)

	if err != nil {
		log.Println("Error while retriving resumes: ", err)
		return nil, err
	}

	for rows.Next() {
		resume := cp.ResumeRes{}

		err := rows.Scan(
			&resume.Id,
			&resume.UserId,
			&resume.Education,
			&resume.Summary,
		)

		if err != nil {
			log.Println("Error while scanning all resumes: ", err)
			return nil, err
		}

		resumes.Resumes = append(resumes.Resumes, &resume)
	}

	log.Println("Successfully fetched all resumes")

	return &resumes, nil
}

func (r *ResumeRepo) Update(req *cp.ResumeUpdateReq) (*cp.ResumeRes, error) {
	resume := cp.ResumeRes{}

	query := `
	UPDATE
		resumes
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.Id == "" || req.Id == "string" {
		return nil, errors.New("nothing to update")
	}
	if req.Education != "" && req.Education != "string" {
		conditions = append(conditions, " education = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Education)
	}
	if req.Summary != "" && req.Summary != "string" {
		conditions = append(conditions, " summary = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Summary)
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
		user_id,
		education,
		summary
	`

	args = append(args, req.Id)

	row := r.db.QueryRow(query, args...)

	err := row.Scan(
		&resume.Id,
		&resume.UserId,
		&resume.Education,
		&resume.Summary,
	)

	if err != nil {
		log.Println("Error while updating resume: ", err)
		return nil, err
	}

	log.Println("Successfully updated resume")

	return &resume, nil
}

func (r *ResumeRepo) Delete(id *cp.Byid) (*cp.Void, error) {
	query := `
	UPDATE 
		resumes
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, id.Id)

	if err != nil {
		log.Println("Error while deleting resume: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("resume with id %s not found", id.Id)
	}

	log.Println("Successfully deleted resume")

	return &cp.Void{
		Success: true,
		Message: "Successfully deleted resume",
	}, nil
}
