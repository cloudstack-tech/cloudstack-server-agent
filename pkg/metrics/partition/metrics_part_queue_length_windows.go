package metrics

import (
	"strings"
	"time"

	"github.com/cloudstack-tech/cloudstack-server-agent/pkg/metrics/collector"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
)

var (
	diskQueueLength = collector.PerformanceQuery.MustAddCounterToQuery("\\LogicalDisk(*)\\Current Disk Queue Length")
)

var _ collector.MetricsCollector = &PartitionQueueLengthCollector{}

type PartitionQueueLengthCollector struct {
	lastCollectTime time.Time
}

func (c *PartitionQueueLengthCollector) init() error {
	collector.PerformanceQuery.CollectData()
	c.lastCollectTime = time.Now()
	return nil
}

func NewPartitionQueueLengthCollector() (*PartitionQueueLengthCollector, error) {
	collector := &PartitionQueueLengthCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

type PartitionQueueLengthMetrics struct {
	Name        string
	QueueLength float64
}

func (c *PartitionQueueLengthCollector) GetName() string {
	return "partition_queue_length"
}

func (c *PartitionQueueLengthCollector) GetValue() (any, error) {
	if time.Since(c.lastCollectTime) < 1000*time.Millisecond {
		time.Sleep(1000*time.Millisecond - time.Since(c.lastCollectTime))
	}
	c.lastCollectTime = time.Now()
	collector.PerformanceQuery.CollectData()
	read, err := collector.PerformanceQuery.GetFormattedCounterArrayDouble(diskQueueLength)
	if err != nil {
		return nil, err
	}

	var diskIoInfo []*PartitionQueueLengthMetrics
	for _, v := range read {
		if strings.Contains(v.Name, "Total") {
			continue
		}
		diskIoInfo = append(diskIoInfo, &PartitionQueueLengthMetrics{
			Name:        v.Name,
			QueueLength: float64(v.Value),
		})
	}

	return diskIoInfo, nil
}

func (c *PartitionQueueLengthCollector) CollectMetrics() (*proto.Metrics, error) {
	return nil, nil
}
