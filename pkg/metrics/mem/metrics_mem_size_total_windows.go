//go:build windows

package metrics

import (
	"fmt"
	"time"

	collector "github.com/cloudstack-tech/cloudstack-server-agent/pkg/metrics/collector"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
	"github.com/shirou/gopsutil/v4/mem"
)

// 接口断言
var _ collector.MetricsCollector = &MemorySizeTotalCollector{}

type MemorySizeTotalCollector struct {
}

func (c *MemorySizeTotalCollector) init() error {
	return nil
}

func NewMemorySizeTotalCollector() (*MemorySizeTotalCollector, error) {
	collector := &MemorySizeTotalCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

func (c *MemorySizeTotalCollector) GetName() string {
	return "memory_size_total"
}

func (c *MemorySizeTotalCollector) GetValue() (any, error) {
	memoryStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	// float64
	return float64(memoryStat.Total) / 1024 / 1024, nil
}

func (c *MemorySizeTotalCollector) CollectMetrics() (*proto.Metrics, error) {
	usage, err := c.GetValue()
	if err != nil {
		return nil, err
	}

	intValue, ok := usage.(float64)
	if !ok {
		return nil, fmt.Errorf("usage is not a int64")
	}

	return &proto.Metrics{
		Name: c.GetName(),
		Value: &proto.Metrics_DoubleValue{
			DoubleValue: intValue * 1024 * 1024,
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Unit:      "MB",
	}, nil
}
