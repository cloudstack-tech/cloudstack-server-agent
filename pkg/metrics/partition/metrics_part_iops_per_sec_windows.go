package metrics

import (
	"strings"
	"time"

	"github.com/cloudstack-tech/cloudstack-server-agent/pkg/metrics/collector"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
)

var (
	diskReadIops  = collector.PerformanceQuery.MustAddCounterToQuery("\\LogicalDisk(*)\\Disk Reads/sec")
	diskWriteIops = collector.PerformanceQuery.MustAddCounterToQuery("\\LogicalDisk(*)\\Disk Writes/sec")
)

var _ collector.MetricsCollector = &PartitionIoPerIopsSecCollector{}

type PartitionIoPerIopsSecCollector struct {
	lastCollectTime time.Time
}

func (c *PartitionIoPerIopsSecCollector) init() error {
	collector.PerformanceQuery.CollectData()
	c.lastCollectTime = time.Now()
	return nil
}

func NewPartitionIoPerIopsSecCollector() (*PartitionIoPerIopsSecCollector, error) {
	collector := &PartitionIoPerIopsSecCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

type PartitionIoPerIopsSecMetrics struct {
	Name      string
	IopsRead  float64
	IopsWrite float64
}

func (c *PartitionIoPerIopsSecCollector) GetName() string {
	return "partition_io_per_iops_sec"
}

func (c *PartitionIoPerIopsSecCollector) GetValue() (any, error) {
	if time.Since(c.lastCollectTime) < 1000*time.Millisecond {
		time.Sleep(1000*time.Millisecond - time.Since(c.lastCollectTime))
	}
	c.lastCollectTime = time.Now()
	collector.PerformanceQuery.CollectData()
	read, err := collector.PerformanceQuery.GetFormattedCounterArrayDouble(diskReadIops)
	if err != nil {
		return nil, err
	}
	write, err := collector.PerformanceQuery.GetFormattedCounterArrayDouble(diskWriteIops)
	if err != nil {
		return nil, err
	}

	var diskIoInfo []*PartitionIoPerIopsSecMetrics
	for k, v := range read {
		if strings.Contains(v.Name, "Total") {
			continue
		}
		diskIoInfo = append(diskIoInfo, &PartitionIoPerIopsSecMetrics{
			Name:      v.Name,
			IopsRead:  float64(v.Value),
			IopsWrite: float64(write[k].Value),
		})
	}

	return diskIoInfo, nil
}

func (c *PartitionIoPerIopsSecCollector) CollectMetrics() (*proto.Metrics, error) {
	return nil, nil
}
