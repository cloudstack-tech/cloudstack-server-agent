package metrics

import (
	"github.com/shirou/gopsutil/v4/cpu"
)

type CpuInfoCollector struct {
}

type CpuInfo struct {
	ModelName string `json:"model_name"`
}

func (c *CpuInfoCollector) init() error {
	return nil
}

func NewCpuInfoCollector() (*CpuInfoCollector, error) {
	collector := &CpuInfoCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

func (c *CpuInfoCollector) GetName() string {
	return "cpu_info"
}

func (c *CpuInfoCollector) CollectMetrics() ([]CpuInfo, error) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	cpuInfos := make([]CpuInfo, len(cpuInfo))
	for i, info := range cpuInfo {
		cpuInfos[i] = CpuInfo{
			ModelName: info.ModelName,
		}
	}

	return cpuInfos, nil
}