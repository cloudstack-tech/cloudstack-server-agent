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
var _ collector.MetricsCollector = &MemorySizeUsedCollector{}

type MemorySizeUsedCollector struct {
}

func (c *MemorySizeUsedCollector) init() error {
	return nil
}

func NewMemorySizeUsedCollector() (*MemorySizeUsedCollector, error) {
	collector := &MemorySizeUsedCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

func (c *MemorySizeUsedCollector) GetName() string {
	return "memory_size_used"
}

func (c *MemorySizeUsedCollector) GetValue() (any, error) {
	memoryStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return float64(memoryStat.Used) / 1024 / 1024, nil
}

func (c *MemorySizeUsedCollector) CollectMetrics() (*proto.Metrics, error) {
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
