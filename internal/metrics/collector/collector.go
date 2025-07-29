package collector

import (
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
)

type MetricsCollector[T any] interface {
	GetName() string
	GetValue() (T, error)
	CollectMetrics() (*proto.Metrics, error)
}
