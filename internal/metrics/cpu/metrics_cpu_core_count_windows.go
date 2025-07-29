//go:build windows

package metrics

import (
	"fmt"
	"strings"
	"time"

	collector "github.com/cloudstack-tech/cloudstack-server-agent/internal/metrics/collector"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
)

var (
	processorCoreCountHandle = performanceQuery.MustAddCounterToQuery("\\Processor Information(*)\\% Processor Utility")
)

var _ collector.MetricsCollector = &CpuCoreCountCollector{}

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

func (c *CpuCoreCountCollector) GetValue() (any, error) {
	performanceQuery.CollectData()

	usage, err := performanceQuery.GetFormattedCounterArrayDouble(processorCoreCountHandle)
	if err != nil {
		return nil, err
	}

	count := 0
	for _, item := range usage {
		if strings.Contains(item.Name, "_Total") {
			continue
		}
		count++
	}

	return count, nil
}

func (c *CpuCoreCountCollector) CollectMetrics() (*proto.Metrics, error) {
	usage, err := c.GetValue()
	if err != nil {
		return nil, err
	}
	floatValue, ok := usage.(int)
	if !ok {
		return nil, fmt.Errorf("usage is not a int")
	}

	return &proto.Metrics{
		Name: c.GetName(),
		Value: &proto.Metrics_Int32Value{
			Int32Value: int32(floatValue),
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Unit:      "core",
	}, nil
}
