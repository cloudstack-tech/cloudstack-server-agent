//go:build windows

package metrics

import (
	"fmt"
	"strings"
	"time"

	collector "github.com/cloudstack-tech/cloudstack-server-agent/internal/metrics/collector"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	processorCoreCountHandle = performanceQuery.MustAddCounterToQuery("\\Processor Information(*)\\% Processor Utility")
)

var _ collector.MetricsCollector[float64] = &CpuCoreCountCollector{}

type CpuCoreCountCollector struct {
}

func (c *CpuCoreCountCollector) init() error {
	performanceQuery.CollectData()
	return nil
}

func NewCpuCoreCountCollector() (*CpuCoreCountCollector, error) {
	collector := &CpuCoreCountCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

func (c *CpuCoreCountCollector) GetName() string {
	return "cpu_core_count"
}

func (c *CpuCoreCountCollector) GetValue() (float64, error) {
	performanceQuery.CollectData()

	usage, err := performanceQuery.GetFormattedCounterArrayDouble(processorCoreCountHandle)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, item := range usage {
		if strings.Contains(item.Name, "_Total") {
			continue
		}
		count++
	}

	return float64(count), nil
}

func (c *CpuCoreCountCollector) CollectMetrics() (*proto.Metrics, error) {
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
		Unit:      "core",
	}, nil
}
