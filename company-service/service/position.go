package service

import (
	cp "company/genprotos"
	st "company/storage/postgres"
	"context"
)

type PositionService struct {
	storage st.Storage
	cp.UnimplementedPositionServiceServer
}

func NewPositionService(storage *st.Storage) *PositionService {
	return &PositionService{
		storage: *storage,
	}
}

func(s *PositionService) Create(ctx context.Context, req *cp.PositionCreateReq) (*cp.PositionRes, error){
	return s.storage.PositionS.Create(req)
}

func(s *PositionService) GetById(ctx context.Context, id *cp.Byid) (*cp.PositionGetByIdRes, error){
	return s.storage.PositionS.GetById(id)
}

func(s *PositionService) GetAll(ctx context.Context, req *cp.PositionGetAllReq) (*cp.PositionGetAllRes, error){
	return s.storage.PositionS.GetAll(req)
}

func(s *PositionService) Update(ctx context.Context, req *cp.PositionUpdateReq) (*cp.PositionRes, error){
	return s.storage.PositionS.Update(req)
}

func(s *PositionService) Delete(ctx context.Context, id *cp.Byid) (*cp.Void, error){
	return s.storage.PositionS.Delete(id)
}
