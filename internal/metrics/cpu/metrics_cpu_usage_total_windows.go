//go:build windows

package metrics

import (
	"fmt"
	"time"

	collector "github.com/cloudstack-tech/cloudstack-server-agent/internal/metrics/collector"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	processorUtilityHandle = performanceQuery.MustAddCounterToQuery("\\Processor Information(_Total)\\% Processor Utility")
)

// 接口断言
var _ collector.MetricsCollector[float64] = &CpuUsageTotalCollector{}

type CpuUsageTotalCollector struct {
}

func (c *CpuUsageTotalCollector) init() error {
	performanceQuery.CollectData()
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

func (c *CpuUsageTotalCollector) GetValue() (float64, error) {
	performanceQuery.CollectData()
	usage, err := performanceQuery.GetFormattedCounterValueDouble(processorUtilityHandle)
	if err != nil {
		return 0, err
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

	value, err := anypb.New(&wrapperspb.DoubleValue{Value: usage})
	if err != nil {
		return nil, fmt.Errorf("failed to create any: %w", err)
	}

	return &proto.Metrics{
		Name:      c.GetName(),
		Value:     value,
		Timestamp: time.Now().Format(time.RFC3339),
		Unit:      "%",
	}, nil
}
