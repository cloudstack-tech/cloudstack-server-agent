package metrics

import (
	"testing"
	"time"
)

func TestGetPartQueueLength(t *testing.T) {
	collector, err := NewPartitionQueueLengthCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	partQueueLengthInfo, err := collector.GetValue()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}
	partQueueLengthInfos, ok := partQueueLengthInfo.([]*PartitionQueueLengthMetrics)
	if !ok {
		t.Fatalf("disk io info is not a []*PartitionQueueLengthMetrics")
	}
	for _, partQueueLength := range partQueueLengthInfos {
		t.Logf("part queue length: %v", partQueueLength)
		t.Logf("part queue length: %v", partQueueLength.QueueLength)
	}
}

func TestGetPartQueueLengthContinuously(t *testing.T) {
	collector, err := NewPartitionQueueLengthCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	count := 0
	for {
		partQueueLengthInfo, err := collector.GetValue()
		if err != nil {
			t.Fatalf("failed to collect metrics: %v", err)
		}
		partQueueLengthInfos, ok := partQueueLengthInfo.([]*PartitionQueueLengthMetrics)
		if !ok {
			t.Fatalf("part queue length info is not a []*PartitionQueueLengthMetrics")
		}
		for _, partQueueLength := range partQueueLengthInfos {
			t.Logf("part queue length: %v", partQueueLength)
		}
		count++
		if count >= 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
