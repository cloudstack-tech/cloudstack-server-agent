package metrics

import (
	"testing"
	"time"
)

func TestGetPartitionIoPerIopsSec(t *testing.T) {
	collector, err := NewPartitionIoPerIopsSecCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	partIoInfo, err := collector.GetValue()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}
	partIoInfos, ok := partIoInfo.([]*PartitionIoPerIopsSecMetrics)
	if !ok {
		t.Fatalf("part io info is not a []*PartitionIoPerIopsSecMetrics")
	}
	for _, partIo := range partIoInfos {
		t.Logf("part io: %v", partIo)
		t.Logf("part io read : %v iops", partIo.IopsRead)
		t.Logf("part io write: %v iops", partIo.IopsWrite)
	}
}

func TestGetPartitionIoPerIopsSecContinuously(t *testing.T) {
	collector, err := NewPartitionIoPerIopsSecCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	count := 0
	for {
		partIoInfo, err := collector.GetValue()
		if err != nil {
			t.Fatalf("failed to collect metrics: %v", err)
		}
		partIoInfos, ok := partIoInfo.([]*PartitionIoPerIopsSecMetrics)
		if !ok {
			t.Fatalf("part io info is not a []*PartitionIoPerIopsSecMetrics")
		}
		for _, partIo := range partIoInfos {
			t.Logf("part io: %v", partIo)
			t.Logf("part io read : %v iops", partIo.IopsRead)
			t.Logf("part io write: %v iops", partIo.IopsWrite)
		}
		count++
		if count >= 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
