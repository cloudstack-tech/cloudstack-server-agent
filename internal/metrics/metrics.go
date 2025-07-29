package metrics

import (
	"github.com/cloudstack-tech/cloudstack-server-agent/internal/metrics/collector"
	cpu "github.com/cloudstack-tech/cloudstack-server-agent/internal/metrics/cpu"
)

var MetricsCollectors = make(map[string]any)

func init() {
	cpuInfoCollector, err := cpu.NewCpuInfoCollector()
	if err != nil {
		panic(err)
	}
	MetricsCollectors["cpu_info"] = cpuInfoCollector

	cpuCoreCountCollector, err := cpu.NewCpuCoreCountCollector()
	if err != nil {
		panic(err)
	}
	MetricsCollectors["cpu_core_count"] = cpuCoreCountCollector

	cpuFrequencyCollector, err := cpu.NewCpuFrequencyCollector()
	if err != nil {
		panic(err)
	}
	MetricsCollectors["cpu_frequency"] = cpuFrequencyCollector
}

func GetMetricsCollector[T any](name string) (collector.MetricsCollector[T], error) {
	c, ok := MetricsCollectors[name]
	if !ok {
		return nil, nil
	}
	collector, ok := c.(collector.MetricsCollector[T])
	if !ok {
		return nil, nil
	}
	return collector, nil
}
