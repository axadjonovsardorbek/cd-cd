package service

import (
	cp "company/genprotos"
	st "company/storage/postgres"
	"context"
)

type ResumeService struct {
	storage st.Storage
	cp.UnimplementedResumeServiceServer
}

func NewResumeService(storage *st.Storage) *ResumeService {
	return &ResumeService{
		storage: *storage,
	}
}

func (s *ResumeService) Create(ctx context.Context, req *cp.ResumeCreateReq) (*cp.ResumeRes, error) {
	return s.storage.ResumeS.Create(req)
}

func (s *ResumeService) GetById(ctx context.Context, id *cp.Byid) (*cp.ResumeGetByIdRes, error) {
	return s.storage.ResumeS.GetById(id)
}

func (s *ResumeService) GetAll(ctx context.Context, req *cp.ResumeGetAllReq) (*cp.ResumeGetAllRes, error) {
	return s.storage.ResumeS.GetAll(req)
}

func (s *ResumeService) Update(ctx context.Context, req *cp.ResumeUpdateReq) (*cp.ResumeRes, error) {
	return s.storage.ResumeS.Update(req)
}

func (s *ResumeService) Delete(ctx context.Context, id *cp.Byid) (*cp.Void, error) {
	return s.storage.ResumeS.Delete(id)
}
