package metrics

import (
	"fmt"
	"time"

	collector "github.com/cloudstack-tech/cloudstack-server-agent/pkg/metrics/collector"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
	"github.com/shirou/gopsutil/v4/cpu"
)

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

	cpuInfos := make([]*proto.CpuInfo, len(cpuInfo))
	for i, info := range cpuInfo {
		cpuInfos[i] = &proto.CpuInfo{
			ModelName: info.ModelName,
		}
	}

	return cpuInfos, nil
}

func (c *CpuInfoCollector) CollectMetrics() (*proto.Metrics, error) {
	cpuInfo, err := c.GetValue()
	if err != nil {
		return nil, fmt.Errorf("failed to get cpu info: %w", err)
	}

	cpuInfoList, ok := cpuInfo.([]*proto.CpuInfo)
	if !ok {
		return nil, fmt.Errorf("cpu info is not a []proto.CpuInfo")
	}

	return &proto.Metrics{
		Name: c.GetName(),
		Value: &proto.Metrics_CpuInfoList{
			CpuInfoList: &proto.CpuInfoList{
				CpuInfos: cpuInfoList,
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Unit:      "cpu_info",
	}, nil
}
