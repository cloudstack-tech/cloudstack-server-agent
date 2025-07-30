package metrics

import (
	"time"

	"github.com/cloudstack-tech/cloudstack-server-agent/pkg/metrics/collector"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
)

var (
	networkTrafficPerPacketsSecMetricsRecv = collector.PerformanceQuery.MustAddCounterToQuery("\\Network Interface(*)\\Packets Received/sec")
	networkTrafficPerPacketsSecMetricsSent = collector.PerformanceQuery.MustAddCounterToQuery("\\Network Interface(*)\\Packets Sent/sec")
)

var _ collector.MetricsCollector = &NetworkTrafficPerPacketsSecCollector{}

type NetworkTrafficPerPacketsSecCollector struct {
	lastCollectTime time.Time
}

func (c *NetworkTrafficPerPacketsSecCollector) init() error {
	collector.PerformanceQuery.CollectData()
	c.lastCollectTime = time.Now()
	return nil
}

func NewNetworkTrafficPerPacketsSecCollector() (*NetworkTrafficPerPacketsSecCollector, error) {
	collector := &NetworkTrafficPerPacketsSecCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

type NetworkTrafficPerPacketsSecMetrics struct {
	Name        string
	PacketsSent float64
	PacketsRecv float64
}

func (c *NetworkTrafficPerPacketsSecCollector) GetName() string {
	return "network_traffic_per_packets_sec"
}

func (c *NetworkTrafficPerPacketsSecCollector) GetValue() (any, error) {
	if time.Since(c.lastCollectTime) < 100*time.Millisecond {
		time.Sleep(100*time.Millisecond - time.Since(c.lastCollectTime))
	}
	c.lastCollectTime = time.Now()
	collector.PerformanceQuery.CollectData()
	recv, err := collector.PerformanceQuery.GetFormattedCounterArrayDouble(networkTrafficPerPacketsSecMetricsRecv)
	if err != nil {
		return nil, err
	}
	sent, err := collector.PerformanceQuery.GetFormattedCounterArrayDouble(networkTrafficPerPacketsSecMetricsSent)
	if err != nil {
		return nil, err
	}

	var networkInfo []*NetworkTrafficPerPacketsSecMetrics
	for k, v := range recv {
		networkInfo = append(networkInfo, &NetworkTrafficPerPacketsSecMetrics{
			Name:        v.Name,
			PacketsRecv: float64(v.Value),
			PacketsSent: float64(sent[k].Value),
		})
	}

	return networkInfo, nil
}

func (c *NetworkTrafficPerPacketsSecCollector) CollectMetrics() (*proto.Metrics, error) {
	return nil, nil
}
