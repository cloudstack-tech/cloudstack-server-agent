package metrics

import (
	"testing"
	"time"
)

func TestGetMemorySizeTotal(t *testing.T) {
	collector, err := NewMemorySizeTotalCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	size, err := collector.GetValue()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}

	t.Logf("memory size total: %.2f MB", size)
}

func TestGetMemorySizeTotalContinuously(t *testing.T) {
	collector, err := NewMemorySizeTotalCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	// 10 次后停止
	count := 0
	for {
		size, err := collector.GetValue()
		if err != nil {
			t.Fatalf("failed to collect metrics: %v", err)
		}

		t.Logf("memory size total: %.2f MB", size)
		count++
		if count >= 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
