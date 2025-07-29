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
var _ collector.MetricsCollector = &MemorySizeFreeCollector{}

type MemorySizeFreeCollector struct {
}

func (c *MemorySizeFreeCollector) init() error {
	return nil
}

func NewMemorySizeFreeCollector() (*MemorySizeFreeCollector, error) {
	collector := &MemorySizeFreeCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

func (c *MemorySizeFreeCollector) GetName() string {
	return "memory_size_free"
}

func (c *MemorySizeFreeCollector) GetValue() (any, error) {
	memoryStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return float64(memoryStat.Free) / 1024 / 1024, nil
}

func (c *MemorySizeFreeCollector) CollectMetrics() (*proto.Metrics, error) {
	usage, err := c.GetValue()
	if err != nil {
		return nil, err
	}

	intValue, ok := usage.(float64)
	if !ok {
		return nil, fmt.Errorf("usage is not a float64")
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
