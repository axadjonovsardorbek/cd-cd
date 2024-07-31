package service

import (
	cp "company/genprotos"
	st "company/storage/postgres"
	"context"
)

type DepartmentService struct {
	storage st.Storage
	cp.UnimplementedDepartmentServiceServer
}

func NewDepartmentService(storage *st.Storage) *DepartmentService {
	return &DepartmentService{
		storage: *storage,
	}
}

func (s *DepartmentService) Create(ctx context.Context, req *cp.DepartmentCreateReq) (*cp.DepartmentRes, error) {
	return s.storage.DepartmentS.Create(req)
}

func (s *DepartmentService) GetById(ctx context.Context, id *cp.Byid) (*cp.DepartmentGetByIdRes, error) {
	return s.storage.DepartmentS.GetById(id)
}

func (s *DepartmentService) GetAll(ctx context.Context, req *cp.DepartmentGetAllReq) (*cp.DepartmentGetAllRes, error) {
	return s.storage.DepartmentS.GetAll(req)
}

func (s *DepartmentService) Update(ctx context.Context, req *cp.DepartmentUpdateReq) (*cp.DepartmentRes, error) {
	return s.storage.DepartmentS.Update(req)
}

func (s *DepartmentService) Delete(ctx context.Context, id *cp.Byid) (*cp.Void, error) {
	return s.storage.DepartmentS.Delete(id)
}
