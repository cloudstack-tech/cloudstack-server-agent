package metrics

import (
	"strings"
	"time"

	"github.com/cloudstack-tech/cloudstack-server-agent/pkg/metrics/collector"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
)

var (
	diskQueueLength = collector.PerformanceQuery.MustAddCounterToQuery("\\PhysicalDisk(*)\\Current Disk Queue Length")
)

var _ collector.MetricsCollector = &DiskQueueLengthCollector{}

type DiskQueueLengthCollector struct {
	lastCollectTime time.Time
}

func (c *DiskQueueLengthCollector) init() error {
	collector.PerformanceQuery.CollectData()
	c.lastCollectTime = time.Now()
	return nil
}

func NewDiskQueueLengthCollector() (*DiskQueueLengthCollector, error) {
	collector := &DiskQueueLengthCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

type DiskQueueLengthMetrics struct {
	Name        string
	QueueLength float64
}

func (c *DiskQueueLengthCollector) GetName() string {
	return "disk_queue_length"
}

func (c *DiskQueueLengthCollector) GetValue() (any, error) {
	if time.Since(c.lastCollectTime) < 1000*time.Millisecond {
		time.Sleep(1000*time.Millisecond - time.Since(c.lastCollectTime))
	}
	c.lastCollectTime = time.Now()
	collector.PerformanceQuery.CollectData()
	read, err := collector.PerformanceQuery.GetFormattedCounterArrayDouble(diskQueueLength)
	if err != nil {
		return nil, err
	}

	var diskIoInfo []*DiskQueueLengthMetrics
	for _, v := range read {
		if strings.Contains(v.Name, "Total") {
			continue
		}
		diskIoInfo = append(diskIoInfo, &DiskQueueLengthMetrics{
			Name:        v.Name,
			QueueLength: float64(v.Value),
		})
	}

	return diskIoInfo, nil
}

func (c *DiskQueueLengthCollector) CollectMetrics() (*proto.Metrics, error) {
	return nil, nil
}
