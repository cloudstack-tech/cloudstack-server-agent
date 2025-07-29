package service

import (
	"context"
	"fmt"

	m "github.com/cloudstack-tech/cloudstack-server-agent/internal/metrics"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
)

type MetricsService struct {
}

func (s *MetricsService) GetMetrics(ctx context.Context, req *proto.MetricsRequest) (*proto.MetricsResponse, error) {
	metrics := make([]*proto.Metrics, 0)
	for _, name := range req.Name {
		collector, err := m.GetMetricsCollector[any](name)
		if err != nil {
			return nil, err
		}
		if collector == nil {
			return nil, fmt.Errorf("collector %s not found", name)
		}
		metric, err := collector.CollectMetrics()
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, metric)
	}

	return &proto.MetricsResponse{
		Metrics: metrics,
	}, nil
}
