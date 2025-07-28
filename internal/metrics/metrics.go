package metrics

import (
	cpu "github.com/cloudstack-tech/cloudstack-server-agent/internal/metrics/cpu"
)

type MetricsCollector[T any] interface {
	GetName() string
	CollectMetrics() (T, error)
}

var metricsCollectors = make(map[string]any)

func init() {
	cpuInfoCollector, err := cpu.NewCpuInfoCollector()
	if err != nil {
		panic(err)
	}
	metricsCollectors["cpu_info"] = cpuInfoCollector
}