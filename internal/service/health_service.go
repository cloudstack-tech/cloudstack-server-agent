package service

import (
	"context"
	"time"

	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
)

type HealthService struct {
	proto.HealthServiceServer
}

const (
	HEALTHY_MESSAGE = "I'm alive!"
)

func (s *HealthService) GetHealth(ctx context.Context, req *proto.HealthRequest) (*proto.HealthResponse, error) {
	return &proto.HealthResponse{
		Status:    proto.HealthStatus_HEALTHY,
		Message:   HEALTHY_MESSAGE,
		Timestamp: time.Now().Format(time.RFC3339),
	}, nil

}
