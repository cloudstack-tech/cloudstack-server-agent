package metrics

import (
	"testing"
	"time"
)

func TestGetMemorySizeUsed(t *testing.T) {
	collector, err := NewMemorySizeUsedCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	size, err := collector.GetValue()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}

	t.Logf("memory size used: %.2f MB", size)
}

func TestGetMemorySizeUsedContinuously(t *testing.T) {
	collector, err := NewMemorySizeUsedCollector()
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

		t.Logf("memory size used: %.2f MB", size)
		count++
		if count >= 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
