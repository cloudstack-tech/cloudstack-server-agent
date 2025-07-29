package metrics

import (
	"time"

	"github.com/cloudstack-tech/cloudstack-server-agent/pkg/metrics/collector"
	"github.com/cloudstack-tech/cloudstack-server-agent/proto"
)

var (
	networkTrafficPerKbitsSecMetricsRecv = collector.PerformanceQuery.MustAddCounterToQuery("\\Network Interface(*)\\Bytes Received/sec")
	networkTrafficPerKbitsSecMetricsSent = collector.PerformanceQuery.MustAddCounterToQuery("\\Network Interface(*)\\Bytes Sent/sec")
)

var _ collector.MetricsCollector = &NetworkTrafficPerKbitsSecCollector{}

type NetworkTrafficPerKbitsSecCollector struct {
	lastCollectTime time.Time
}

func (c *NetworkTrafficPerKbitsSecCollector) init() error {
	collector.PerformanceQuery.CollectData()
	c.lastCollectTime = time.Now()
	return nil
}

func NewNetworkTrafficPerKbitsSecCollector() (*NetworkTrafficPerKbitsSecCollector, error) {
	collector := &NetworkTrafficPerKbitsSecCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

type NetworkTrafficPerKbitsSecMetrics struct {
	Name      string
	KbitsSent float64
	KbitsRecv float64
}

func (c *NetworkTrafficPerKbitsSecCollector) GetName() string {
	return "network_traffic_per_kbits_sec"
}

func (c *NetworkTrafficPerKbitsSecCollector) GetValue() (any, error) {
	if time.Since(c.lastCollectTime) < 100*time.Millisecond {
		time.Sleep(100*time.Millisecond - time.Since(c.lastCollectTime))
	}
	c.lastCollectTime = time.Now()
	collector.PerformanceQuery.CollectData()
	recv, err := collector.PerformanceQuery.GetFormattedCounterArrayDouble(networkTrafficPerKbitsSecMetricsRecv)
	if err != nil {
		return nil, err
	}
	sent, err := collector.PerformanceQuery.GetFormattedCounterArrayDouble(networkTrafficPerKbitsSecMetricsSent)
	if err != nil {
		return nil, err
	}

	var networkInfo []*NetworkTrafficPerKbitsSecMetrics
	for k, v := range recv {
		networkInfo = append(networkInfo, &NetworkTrafficPerKbitsSecMetrics{
			Name:      v.Name,
			KbitsRecv: float64(v.Value) * 8 / 1024.0,
			KbitsSent: float64(sent[k].Value) * 8 / 1024.0,
		})
	}

	return networkInfo, nil
}

func (c *NetworkTrafficPerKbitsSecCollector) CollectMetrics() (*proto.Metrics, error) {
	return nil, nil
}
