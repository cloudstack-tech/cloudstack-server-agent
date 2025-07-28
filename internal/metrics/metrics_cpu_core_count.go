//go:build windows

package metrics

import "strings"

var (
	processorCoreCountHandle = performanceQuery.MustAddCounterToQuery("\\Processor Information(*)\\% Processor Utility")
)

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

func (c *CpuCoreCountCollector) CollectMetrics() (float64, error) {
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