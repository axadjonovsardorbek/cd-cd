package clients

import (
	"gateway/config"
	cp "gateway/genprotos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	Company    cp.CompanyServiceClient
	Department cp.DepartmentServiceClient
	Position   cp.PositionServiceClient
	Resume     cp.ResumeServiceClient
}

func NewGrpcClients(cfg *config.Config) (*GrpcClients, error) {
	connC, err := grpc.NewClient("company-service"+cfg.COMPANY_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GrpcClients{
		Company:    cp.NewCompanyServiceClient(connC),
		Department: cp.NewDepartmentServiceClient(connC),
		Position:   cp.NewPositionServiceClient(connC),
		Resume:     cp.NewResumeServiceClient(connC),
	}, nil
}
