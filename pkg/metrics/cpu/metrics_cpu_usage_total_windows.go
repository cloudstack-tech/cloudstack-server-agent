//go:build windows

package metrics

import (
	"fmt"
	"time"

	collector "github.com/cloudstack-tech/cloudstack-server-agent/pkg/metrics/collector"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
)

var (
	processorUtilityHandle = collector.PerformanceQuery.MustAddCounterToQuery("\\Processor Information(_Total)\\% Processor Utility")
)

// 接口断言
var _ collector.MetricsCollector = &CpuUsageTotalCollector{}

type CpuUsageTotalCollector struct {
}

func (c *CpuUsageTotalCollector) init() error {
	collector.PerformanceQuery.CollectData()
	return nil
}

func NewCpuUsageTotalCollector() (*CpuUsageTotalCollector, error) {
	collector := &CpuUsageTotalCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

func (c *CpuUsageTotalCollector) GetName() string {
	return "cpu_usage_total"
}

func (c *CpuUsageTotalCollector) GetValue() (any, error) {
	collector.PerformanceQuery.CollectData()
	usage, err := collector.PerformanceQuery.GetFormattedCounterValueDouble(processorUtilityHandle)
	if err != nil {
		return nil, err
	}
	if usage > 100 {
		usage = 100
	}

	return usage, nil
}

func (c *CpuUsageTotalCollector) CollectMetrics() (*proto.Metrics, error) {
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
