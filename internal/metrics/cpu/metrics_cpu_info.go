package metrics

import (
	"fmt"
	"time"

	collector "github.com/cloudstack-tech/cloudstack-server-agent/internal/metrics/collector"
	pb "github.com/cloudstack-tech/cloudstack-server-agent/proto"
	"github.com/shirou/gopsutil/v4/cpu"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

type CpuInfo struct {
	ModelName string
}

var _ collector.MetricsCollector[[]CpuInfo] = &CpuInfoCollector{}

type CpuInfoCollector struct {
}

func NewCpuInfoCollector() (*CpuInfoCollector, error) {
	collector := &CpuInfoCollector{}
	return collector, nil
}

func (c *CpuInfoCollector) GetName() string {
	return "cpu_info"
}

func (c *CpuInfoCollector) GetValue() ([]CpuInfo, error) {

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

func (c *CpuInfoCollector) CollectMetrics() (*pb.Metrics, error) {
	cpuInfo, err := c.GetValue()
	if err != nil {
		return nil, fmt.Errorf("failed to get cpu info: %w", err)
	}

	cpuInfoList := make([]any, len(cpuInfo))
	for i, info := range cpuInfo {
		cpuInfoList[i] = info
	}

	cpuInfopb, err := structpb.NewList(cpuInfoList)
	if err != nil {
		return nil, fmt.Errorf("failed to create struct: %w", err)
	}

	value, err := anypb.New(cpuInfopb)
	if err != nil {
		return nil, fmt.Errorf("failed to create any: %w", err)
	}

	return &pb.Metrics{
		Name:      c.GetName(),
		Value:     value,
		Timestamp: time.Now().Format(time.RFC3339),
		Unit:      "",
	}, nil
}
