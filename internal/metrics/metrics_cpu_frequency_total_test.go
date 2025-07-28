package metrics

import (
	"testing"
	"time"
)

func TestGetCpuFrequency(t *testing.T) {
	collector, err := NewCpuFrequencyCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	usage, err := collector.CollectMetrics()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}

	t.Logf("cpu frequency: %f", usage)
}

func TestGetCpuFrequencyContinuously(t *testing.T) {
	collector, err := NewCpuFrequencyCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	// 10 次后停止
	count := 0
	for {
		usage, err := collector.CollectMetrics()
		if err != nil {
			t.Fatalf("failed to collect metrics: %v", err)
		}

		t.Logf("cpu frequency: %f", usage)
		count++
		if count >= 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}