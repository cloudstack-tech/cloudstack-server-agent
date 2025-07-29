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
var _ collector.MetricsCollector = &MemoryUsageTotalCollector{}

type MemoryUsageTotalCollector struct {
}

func (c *MemoryUsageTotalCollector) init() error {
	return nil
}

func NewMemoryUsageTotalCollector() (*MemoryUsageTotalCollector, error) {
	collector := &MemoryUsageTotalCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

func (c *MemoryUsageTotalCollector) GetName() string {
	return "memory_usage_total"
}

func (c *MemoryUsageTotalCollector) GetValue() (any, error) {
	memoryStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return memoryStat.UsedPercent, nil
}

func (c *MemoryUsageTotalCollector) CollectMetrics() (*proto.Metrics, error) {
	usage, err := c.GetValue()
	if err != nil {
		return nil, err
	}

	floatValue, ok := usage.(float64)
	if !ok {
		return nil, fmt.Errorf("usage is not a float64")
	}

	return &proto.Metrics{
		Name: c.GetName(),
		Value: &proto.Metrics_DoubleValue{
			DoubleValue: floatValue,
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Unit:      "%",
	}, nil
}
