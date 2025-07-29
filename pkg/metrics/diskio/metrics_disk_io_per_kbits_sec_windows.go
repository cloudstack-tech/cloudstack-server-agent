package metrics

import (
	"strings"
	"time"

	"github.com/cloudstack-tech/cloudstack-server-agent/pkg/metrics/collector"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
)

var (
	diskReadBytes  = collector.PerformanceQuery.MustAddCounterToQuery("\\PhysicalDisk(*)\\Disk Read Bytes/sec")
	diskWriteBytes = collector.PerformanceQuery.MustAddCounterToQuery("\\PhysicalDisk(*)\\Disk Write Bytes/sec")
)

var _ collector.MetricsCollector = &DiskIoPerKbitsSecCollector{}

type DiskIoPerKbitsSecCollector struct {
	lastCollectTime time.Time
}

func (c *DiskIoPerKbitsSecCollector) init() error {
	collector.PerformanceQuery.CollectData()
	c.lastCollectTime = time.Now()
	return nil
}

func NewDiskIoPerKbitsSecCollector() (*DiskIoPerKbitsSecCollector, error) {
	collector := &DiskIoPerKbitsSecCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

type DiskIoPerKbitsSecMetrics struct {
	Name       string
	KbitsRead  float64
	KbitsWrite float64
}

func (c *DiskIoPerKbitsSecCollector) GetName() string {
	return "disk_io_per_kbits_sec"
}

func (c *DiskIoPerKbitsSecCollector) GetValue() (any, error) {
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

	var diskIoInfo []*DiskIoPerKbitsSecMetrics
	for k, v := range read {
		if strings.Contains(v.Name, "Total") {
			continue
		}
		// Byte/s to Kbits/s
		diskIoInfo = append(diskIoInfo, &DiskIoPerKbitsSecMetrics{
			Name:       v.Name,
			KbitsRead:  float64(v.Value) * 8 / 1000.0,
			KbitsWrite: float64(write[k].Value) * 8 / 1000.0,
		})
	}

	return diskIoInfo, nil
}

func (c *DiskIoPerKbitsSecCollector) CollectMetrics() (*proto.Metrics, error) {
	return nil, nil
}
