package metrics

import (
	"fmt"

	collector "github.com/cloudstack-tech/cloudstack-server-agent/internal/metrics/collector"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
	"github.com/shirou/gopsutil/v4/cpu"
)

type CpuInfo struct {
	ModelName string
}

var _ collector.MetricsCollector = &CpuInfoCollector{}

type CpuInfoCollector struct {
}

func NewCpuInfoCollector() (*CpuInfoCollector, error) {
	collector := &CpuInfoCollector{}
	return collector, nil
}

func (c *CpuInfoCollector) GetName() string {
	return "cpu_info"
}

func (c *CpuInfoCollector) GetValue() (any, error) {

	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, fmt.Errorf("failed to get cpu info: %w", err)
	}

	cpuInfos := make([]CpuInfo, len(cpuInfo))
	for i, info := range cpuInfo {
		cpuInfos[i] = CpuInfo{
			ModelName: info.ModelName,
		}
	}

	return cpuInfos, nil
}

func (c *CpuInfoCollector) CollectMetrics() (*proto.Metrics, error) {
	// cpuInfo, err := c.GetValue()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get cpu info: %w", err)
	// }

	return nil, nil
}
