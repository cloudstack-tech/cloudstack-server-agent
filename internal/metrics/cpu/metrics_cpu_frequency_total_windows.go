//go:build windows

package metrics

import (
	"fmt"
	"time"

	collector "github.com/cloudstack-tech/cloudstack-server-agent/internal/metrics/collector"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
)

var (
	processorFrequencyHandle = performanceQuery.MustAddCounterToQuery("\\Processor Information(_Total)\\Actual Frequency")
)

var _ collector.MetricsCollector = &CpuFrequencyCollector{}

type CpuFrequencyCollector struct {
}

func (c *CpuFrequencyCollector) init() error {
	performanceQuery.CollectData()
	return nil
}

func NewCpuFrequencyCollector() (*CpuFrequencyCollector, error) {
	collector := &CpuFrequencyCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

func (c *CpuFrequencyCollector) GetName() string {
	return "cpu_frequency_total"
}

func (c *CpuFrequencyCollector) GetValue() (any, error) {
	usage, err := performanceQuery.GetFormattedCounterValueDouble(processorFrequencyHandle)
	if err != nil {
		return nil, err
	}

	return usage / 1000, nil
}

func (c *CpuFrequencyCollector) CollectMetrics() (*proto.Metrics, error) {
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
		Unit:      "MHz",
	}, nil
}
