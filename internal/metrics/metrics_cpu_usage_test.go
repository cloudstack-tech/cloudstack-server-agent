package metrics

import (
	"testing"
	"time"
)

func TestGetCpuUsage(t *testing.T) {
	collector := NewCpuUsageCollector()
	usage, err := collector.CollectMetrics()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}

	t.Logf("cpu usage: %f", usage)
}

func TestGetCpuUsageContinuously(t *testing.T) {
	collector := NewCpuUsageCollector()
	// 10 次后停止
	count := 0
	for {
		usage, err := collector.CollectMetrics()
		if err != nil {
			t.Fatalf("failed to collect metrics: %v", err)
		}

		t.Logf("cpu usage: %f", usage)
		count++
		if count >= 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}