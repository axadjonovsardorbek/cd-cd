package postgres

import (
	"company/config"
	"company/storage"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db          *sql.DB
	CompanyS    storage.CompanyI
	DepartmentS storage.DepartmentI
	PositionS   storage.PositionI
	ResumeS     storage.ResumeI
}

func NewPostgresStorage(config config.Config) (*Storage, error) {
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	company := NewCompanyRepo(db)
	department := NewDepartmentRepo(db)
	position := NewPositionRepo(db)
	resume := NewResumeRepo(db)

	return &Storage{
		Db:          db,
		CompanyS:    company,
		DepartmentS: department,
		PositionS:   position,
		ResumeS:     resume,
	}, nil
}
