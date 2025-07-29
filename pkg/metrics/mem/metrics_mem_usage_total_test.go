package metrics

import (
	"testing"
	"time"
)

func TestGetMemoryUsage(t *testing.T) {
	collector, err := NewMemoryUsageTotalCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	usage, err := collector.GetValue()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}

	t.Logf("memory usage: %f", usage)
}

func TestGetMemoryUsageContinuously(t *testing.T) {
	collector, err := NewMemoryUsageTotalCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	// 10 次后停止
	count := 0
	for {
		usage, err := collector.GetValue()
		if err != nil {
			t.Fatalf("failed to collect metrics: %v", err)
		}

		t.Logf("memory usage: %f", usage)
		count++
		if count >= 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
