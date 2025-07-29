//go:build windows

package metrics

import (
	"testing"
	"time"
)

func TestGetCpuCoreCount(t *testing.T) {
	collector, err := NewCpuCoreCountCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	cpuCoreCount, err := collector.GetValue()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}
	t.Logf("cpu core count: %d", cpuCoreCount)
}

func TestGetCpuCoreCountContinuously(t *testing.T) {
	collector, err := NewCpuCoreCountCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	cpuCoreCount, err := collector.GetValue()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}
	t.Logf("cpu core count: %d", cpuCoreCount)
	count := 0
	for {
		cpuCoreCount, err := collector.GetValue()
		if err != nil {
			t.Fatalf("failed to collect metrics: %v", err)
		}
		t.Logf("cpu core count: %d", cpuCoreCount)
		count++
		if count >= 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
