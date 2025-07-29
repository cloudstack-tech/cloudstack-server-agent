package metrics

import (
	"testing"
	"time"
)

func TestGetMemorySizeFree(t *testing.T) {
	collector, err := NewMemorySizeFreeCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	size, err := collector.GetValue()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}

	t.Logf("memory size free: %.2f MB", size)
}

func TestGetMemorySizeFreeContinuously(t *testing.T) {
	collector, err := NewMemorySizeFreeCollector()
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

		t.Logf("memory size free: %.2f MB", size)
		count++
		if count >= 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
