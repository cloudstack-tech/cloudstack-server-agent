package server

import (
	"log"
	"net"

	"github.com/cloudstack-tech/cloudstack-server-agent/internal/service"
	pb "github.com/cloudstack-tech/cloudstack-server-agent/proto"
	"google.golang.org/grpc"
)

const (
	Address string = ":8000"
	Network string = "tcp"
)

func RunGrpcServer() {
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	grpcServer := grpc.NewServer()
	pb.RegisterMetricsServiceServer(
		grpcServer, &service.MetricsService{},
	)
	pb.RegisterHealthServiceServer(
		grpcServer, &service.HealthService{},
	)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
