package main

import (
	cf "company/config"
	cp "company/genprotos"
	service "company/service"
	"company/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config := cf.Load()

	db, err := postgres.NewPostgresStorage(config)

	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", config.COMPANY_SERVICE_PORT)

	if err != nil {
		log.Fatalf("Failed to listen tcp: %v", err)
	}

	s := grpc.NewServer()

	cp.RegisterCompanyServiceServer(s, service.NewCompanyService(db))
	cp.RegisterDepartmentServiceServer(s, service.NewDepartmentService(db))
	cp.RegisterPositionServiceServer(s, service.NewPositionService(db))
	cp.RegisterResumeServiceServer(s, service.NewResumeService(db))

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}