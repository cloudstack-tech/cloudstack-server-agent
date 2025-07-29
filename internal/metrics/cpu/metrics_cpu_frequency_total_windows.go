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
	processorFrequencyHandle = performanceQuery.MustAddCounterToQuery("\\Processor Information(_Total)\\Actual Frequency")
)

var _ collector.MetricsCollector[float64] = &CpuFrequencyCollector{}

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

func (c *CpuFrequencyCollector) GetValue() (float64, error) {
	usage, err := performanceQuery.GetFormattedCounterValueDouble(processorFrequencyHandle)
	if err != nil {
		return 0, err
	}

	return usage / 1000, nil
}

func (c *CpuFrequencyCollector) CollectMetrics() (*proto.Metrics, error) {
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
		Unit:      "MHz",
	}, nil
}
