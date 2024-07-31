package storage

import cp "company/genprotos"

type CompanyI interface {
	Create(*cp.CompanyCreateReq) (*cp.CompanyRes, error)
	GetById(*cp.Byid) (*cp.CompanyGetByIdRes, error)
	GetAll(*cp.CompanyGetAllReq) (*cp.CompanyGetAllRes, error)
	Update(*cp.CompanyUpdateReq) (*cp.CompanyRes, error)
	Delete(*cp.Byid) (*cp.Void, error)
}

type DepartmentI interface {
	Create(*cp.DepartmentCreateReq) (*cp.DepartmentRes, error)
	GetById(*cp.Byid) (*cp.DepartmentGetByIdRes, error)
	GetAll(*cp.DepartmentGetAllReq) (*cp.DepartmentGetAllRes, error)
	Update(*cp.DepartmentUpdateReq) (*cp.DepartmentRes, error)
	Delete(*cp.Byid) (*cp.Void, error)
}

type PositionI interface {
	Create(*cp.PositionCreateReq) (*cp.PositionRes, error)
	GetById(*cp.Byid) (*cp.PositionGetByIdRes, error)
	GetAll(*cp.PositionGetAllReq) (*cp.PositionGetAllRes, error)
	Update(*cp.PositionUpdateReq) (*cp.PositionRes, error)
	Delete(*cp.Byid) (*cp.Void, error)
}

type ResumeI interface {
	Create(*cp.ResumeCreateReq) (*cp.ResumeRes, error)
	GetById(*cp.Byid) (*cp.ResumeGetByIdRes, error)
	GetAll(*cp.ResumeGetAllReq) (*cp.ResumeGetAllRes, error)
	Update(*cp.ResumeUpdateReq) (*cp.ResumeRes, error)
	Delete(*cp.Byid) (*cp.Void, error)
}
