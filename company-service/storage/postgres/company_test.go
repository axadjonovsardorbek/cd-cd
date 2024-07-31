package postgres_test

import (
	"testing"
	// "database/sql"
	"regexp"
	// "time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	cp "company/genprotos"
	"company/storage/postgres"
)

func TestCreateCompany(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewCompanyRepo(db)

	req := &cp.CompanyCreateReq{
		Name:     "Test Company",
		Address:  "123 Test St",
		Industry: "Testing",
		Website:  "www.test.com",
	}

	id := uuid.New().String()

	rows := sqlmock.NewRows([]string{"id", "name", "address", "industry", "website"}).
		AddRow(id, req.Name, req.Address, req.Industry, req.Website)

	query := regexp.QuoteMeta(`
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
	`)

	mock.ExpectQuery(query).
		WithArgs(sqlmock.AnyArg(), req.Name, req.Address, req.Industry, req.Website).
		WillReturnRows(rows)

	company, err := repo.Create(req)
	assert.NoError(t, err)
	assert.NotNil(t, company)
	assert.Equal(t, company.Name, req.Name)
}

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewCompanyRepo(db)

	id := uuid.New().String()
	expectedCompany := &cp.CompanyRes{
		Id:       id,
		Name:     "Test Company",
		Address:  "123 Test St",
		Industry: "Testing",
		Website:  "www.test.com",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "address", "industry", "website"}).
		AddRow(expectedCompany.Id, expectedCompany.Name, expectedCompany.Address, expectedCompany.Industry, expectedCompany.Website)

	query := regexp.QuoteMeta(`
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
	`)

	mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	company, err := repo.GetById(&cp.Byid{Id: id})
	assert.NoError(t, err)
	assert.NotNil(t, company)
	assert.Equal(t, expectedCompany.Name, company.Company.Name)
}

func TestGetAllCompanies(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewCompanyRepo(db)

	req := &cp.CompanyGetAllReq{
		Name: "string",
		Address: "string",
		Industry: "string",
		Website: "string",
		Filter: &cp.Pagination{
			Limit:  10,
			Offset: 0,
		},
	}

	expectedCompanies := []*cp.CompanyRes{
		{
			Id:       "123e4567-e89b-12d3-a456-426614174000",
			Name:     "TechCorp",
			Address:  "123 Tech Street",
			Industry: "Technology",
			Website:  "www.techcorp.com",
		},
		{
			Id:       "223e4567-e89b-12d3-a456-426614174001",
			Name:     "HealthCorp",
			Address:  "456 Health Avenue",
			Industry: "Healthcare",
			Website:  "www.healthcorp.com",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "address", "industry", "website"})
	for _, company := range expectedCompanies {
		rows.AddRow(company.Id, company.Name, company.Address, company.Industry, company.Website)
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
		deleted_at = 0
	LIMIT $1 OFFSET $2`

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(req.Filter.Limit, req.Filter.Offset).WillReturnRows(rows)

	companies, err := repo.GetAll(req)
	assert.NoError(t, err)
	assert.NotNil(t, companies)
	assert.Len(t, companies.Companies, len(expectedCompanies))

	for i, company := range companies.Companies {
		assert.Equal(t, expectedCompanies[i].Id, company.Id)
		assert.Equal(t, expectedCompanies[i].Name, company.Name)
		assert.Equal(t, expectedCompanies[i].Address, company.Address)
		assert.Equal(t, expectedCompanies[i].Industry, company.Industry)
		assert.Equal(t, expectedCompanies[i].Website, company.Website)
	}
}


func TestUpdateCompany(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewCompanyRepo(db)

	req := &cp.CompanyUpdateReq{
		Id:       uuid.New().String(),
		Name:     "Updated Test Company",
		Address:  "123 Updated St",
		Industry: "Updated Testing",
		Website:  "www.updatedtest.com",
	}

	expectedCompany := &cp.CompanyRes{
		Id:       req.Id,
		Name:     req.Name,
		Address:  req.Address,
		Industry: req.Industry,
		Website:  req.Website,
	}

	query := regexp.QuoteMeta(`
	UPDATE
		companies
	SET 
		 name = $1, address = $2, industry = $3, website = $4, updated_at = now() WHERE id = $5 AND deleted_at = 0 
	RETURNING
		id,
		name,
		address,
		industry,
		website
	`)

	rows := sqlmock.NewRows([]string{"id", "name", "address", "industry", "website"}).
		AddRow(expectedCompany.Id, expectedCompany.Name, expectedCompany.Address, expectedCompany.Industry, expectedCompany.Website)

	mock.ExpectQuery(query).WithArgs(req.Name, req.Address, req.Industry, req.Website, req.Id).WillReturnRows(rows)

	company, err := repo.Update(req)
	assert.NoError(t, err)
	assert.NotNil(t, company)
	assert.Equal(t, expectedCompany.Name, company.Name)
}

func TestDeleteCompany(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgres.NewCompanyRepo(db)

	id := uuid.New().String()

	query := regexp.QuoteMeta(`
	UPDATE 
		companies
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`)

	mock.ExpectExec(query).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = repo.Delete(&cp.Byid{Id: id})
	assert.NoError(t, err)
}
