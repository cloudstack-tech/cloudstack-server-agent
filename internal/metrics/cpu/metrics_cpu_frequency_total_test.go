//go:build windows

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
	value, err := collector.GetValue()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}

	t.Logf("cpu frequency: %f", value)
}

func TestGetCpuFrequencyContinuously(t *testing.T) {
	collector, err := NewCpuFrequencyCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	// 10 次后停止
	count := 0
	for {
		value, err := collector.GetValue()
		if err != nil {
			t.Fatalf("failed to collect metrics: %v", err)
		}

		t.Logf("cpu frequency: %f", value)
		count++
		if count >= 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
