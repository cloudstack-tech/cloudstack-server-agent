package metrics

import (
	"strings"
	"time"

	"github.com/cloudstack-tech/cloudstack-server-agent/pkg/metrics/collector"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
)

var (
	diskReadBytes  = collector.PerformanceQuery.MustAddCounterToQuery("\\LogicalDisk(*)\\Disk Read Bytes/sec")
	diskWriteBytes = collector.PerformanceQuery.MustAddCounterToQuery("\\LogicalDisk(*)\\Disk Write Bytes/sec")
)

var _ collector.MetricsCollector = &PartitionIoPerKbitsSecCollector{}

type PartitionIoPerKbitsSecCollector struct {
	lastCollectTime time.Time
}

func (c *PartitionIoPerKbitsSecCollector) init() error {
	collector.PerformanceQuery.CollectData()
	c.lastCollectTime = time.Now()
	return nil
}

func NewPartitionIoPerKbitsSecCollector() (*PartitionIoPerKbitsSecCollector, error) {
	collector := &PartitionIoPerKbitsSecCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

type PartitionIoPerKbitsSecMetrics struct {
	Name       string
	KbitsRead  float64
	KbitsWrite float64
}

func (c *PartitionIoPerKbitsSecCollector) GetName() string {
	return "partition_io_per_kbits_sec"
}

func (c *PartitionIoPerKbitsSecCollector) GetValue() (any, error) {
	if time.Since(c.lastCollectTime) < 1000*time.Millisecond {
		time.Sleep(1000*time.Millisecond - time.Since(c.lastCollectTime))
	}
	c.lastCollectTime = time.Now()
	collector.PerformanceQuery.CollectData()
	read, err := collector.PerformanceQuery.GetFormattedCounterArrayDouble(diskReadBytes)
	if err != nil {
		return nil, err
	}
	write, err := collector.PerformanceQuery.GetFormattedCounterArrayDouble(diskWriteBytes)
	if err != nil {
		return nil, err
	}

	var diskIoInfo []*PartitionIoPerKbitsSecMetrics
	for k, v := range read {
		if strings.Contains(v.Name, "Total") {
			continue
		}
		// Byte/s to Kbits/s
		diskIoInfo = append(diskIoInfo, &PartitionIoPerKbitsSecMetrics{
			Name:       v.Name,
			KbitsRead:  float64(v.Value) * 8 / 1000.0,
			KbitsWrite: float64(write[k].Value) * 8 / 1000.0,
		})
	}

	return diskIoInfo, nil
}

func (c *PartitionIoPerKbitsSecCollector) CollectMetrics() (*proto.Metrics, error) {
	return nil, nil
}
