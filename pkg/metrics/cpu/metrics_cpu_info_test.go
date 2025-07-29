package metrics

import "testing"

func TestGetCpuInfo(t *testing.T) {
	collector, err := NewCpuInfoCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	cpuInfo, err := collector.GetValue()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}
	t.Logf("cpu info: %v", cpuInfo)
}
