package collector

import (
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
)

type MetricsCollector interface {
	GetName() string
	GetValue() (any, error)
	CollectMetrics() (*proto.Metrics, error)
}
