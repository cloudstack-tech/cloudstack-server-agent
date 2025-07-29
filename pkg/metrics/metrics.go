package metrics

import (
	collector "github.com/cloudstack-tech/cloudstack-server-agent/pkg/metrics/collector"
	cpu "github.com/cloudstack-tech/cloudstack-server-agent/pkg/metrics/cpu"
	network "github.com/cloudstack-tech/cloudstack-server-agent/pkg/metrics/network"
)

var MetricsCollectors = make(map[string]collector.MetricsCollector)

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

	cpuUsageCollector, err := cpu.NewCpuUsageTotalCollector()
	if err != nil {
		panic(err)
	}
	MetricsCollectors["cpu_usage_total"] = cpuUsageCollector

	networkTrafficPerKbitsSecCollector, err := network.NewNetworkTrafficPerKbitsSecCollector()
	if err != nil {
		panic(err)
	}
	MetricsCollectors["network_traffic_per_kbits_sec"] = networkTrafficPerKbitsSecCollector
}

func GetMetricsCollector(name string) (collector.MetricsCollector, error) {
	c, ok := MetricsCollectors[name]
	if !ok {
		return nil, nil
	}
	return c, nil
}
