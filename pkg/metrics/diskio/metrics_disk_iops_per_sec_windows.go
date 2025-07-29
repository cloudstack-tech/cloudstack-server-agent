package metrics

import (
	"strings"
	"time"

	"github.com/cloudstack-tech/cloudstack-server-agent/pkg/metrics/collector"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
)

var (
	diskReadIops  = collector.PerformanceQuery.MustAddCounterToQuery("\\PhysicalDisk(*)\\Disk Reads/sec")
	diskWriteIops = collector.PerformanceQuery.MustAddCounterToQuery("\\PhysicalDisk(*)\\Disk Writes/sec")
)

var _ collector.MetricsCollector = &DiskIoPerIopsSecCollector{}

type DiskIoPerIopsSecCollector struct {
	lastCollectTime time.Time
}

func (c *DiskIoPerIopsSecCollector) init() error {
	collector.PerformanceQuery.CollectData()
	c.lastCollectTime = time.Now()
	return nil
}

func NewDiskIoPerIopsSecCollector() (*DiskIoPerIopsSecCollector, error) {
	collector := &DiskIoPerIopsSecCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

type DiskIoPerIopsSecMetrics struct {
	Name      string
	IopsRead  float64
	IopsWrite float64
}

func (c *DiskIoPerIopsSecCollector) GetName() string {
	return "disk_io_per_iops_sec"
}

func (c *DiskIoPerIopsSecCollector) GetValue() (any, error) {
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

	var diskIoInfo []*DiskIoPerIopsSecMetrics
	for k, v := range read {
		if strings.Contains(v.Name, "Total") {
			continue
		}
		diskIoInfo = append(diskIoInfo, &DiskIoPerIopsSecMetrics{
			Name:      v.Name,
			IopsRead:  float64(v.Value),
			IopsWrite: float64(write[k].Value),
		})
	}

	return diskIoInfo, nil
}

func (c *DiskIoPerIopsSecCollector) CollectMetrics() (*proto.Metrics, error) {
	return nil, nil
}
