package metrics

import (
	"testing"
	"time"
)

func TestGetNetworkTrafficPerPacketsSec(t *testing.T) {
	collector, err := NewNetworkTrafficPerPacketsSecCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	networkInfo, err := collector.GetValue()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}
	networkInfos, ok := networkInfo.([]*NetworkTrafficPerPacketsSecMetrics)
	if !ok {
		t.Fatalf("network info is not a []*NetworkTrafficPerPacketsSecMetrics")
	}
	for _, network := range networkInfos {
		t.Logf("network: %v", network)
	}
}

func TestGetNetworkTrafficPerPacketsSecContinuously(t *testing.T) {
	collector, err := NewNetworkTrafficPerPacketsSecCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	count := 0
	for {
		networkInfo, err := collector.GetValue()
		if err != nil {
			t.Fatalf("failed to collect metrics: %v", err)
		}
		networkInfos, ok := networkInfo.([]*NetworkTrafficPerPacketsSecMetrics)
		if !ok {
			t.Fatalf("network info is not a []*NetworkTrafficPerPacketsSecMetrics")
		}
		for _, network := range networkInfos {
			t.Logf("network: %v", network)
		}
		count++
		if count >= 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
