package service

import (
	cp "company/genprotos"
	st "company/storage/postgres"
	"context"
	"fmt"
)

type CompanyService struct {
	storage st.Storage
	cp.UnimplementedCompanyServiceServer
}

func NewCompanyService(storage *st.Storage) *CompanyService {
	return &CompanyService{
		storage: *storage,
	}
}

func (s *CompanyService) Create(ctx context.Context, req *cp.CompanyCreateReq) (*cp.CompanyRes, error) {
	return s.storage.CompanyS.Create(req)
}

func (s *CompanyService) GetById(ctx context.Context, id *cp.Byid) (*cp.CompanyGetByIdRes, error) {
	return s.storage.CompanyS.GetById(id)
}

func (s *CompanyService) GetAll(ctx context.Context, req *cp.CompanyGetAllReq) (*cp.CompanyGetAllRes, error) {
	fmt.Println(req.Address)
	fmt.Println(req.Name)
	fmt.Println(req.Industry)
	fmt.Println(req.Website)
	return s.storage.CompanyS.GetAll(req)
}

func (s *CompanyService) Update(ctx context.Context, req *cp.CompanyUpdateReq) (*cp.CompanyRes, error) {
	return s.storage.CompanyS.Update(req)
}

func (s *CompanyService) Delete(ctx context.Context, id *cp.Byid) (*cp.Void, error) {
	return s.storage.CompanyS.Delete(id)
}
