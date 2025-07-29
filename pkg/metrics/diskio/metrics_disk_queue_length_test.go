package metrics

import (
	"testing"
	"time"
)

func TestGetDiskQueueLength(t *testing.T) {
	collector, err := NewDiskQueueLengthCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	diskIoInfo, err := collector.GetValue()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}
	diskIoInfos, ok := diskIoInfo.([]*DiskQueueLengthMetrics)
	if !ok {
		t.Fatalf("disk io info is not a []*DiskQueueLengthMetrics")
	}
	for _, diskIo := range diskIoInfos {
		t.Logf("disk io: %v", diskIo)
		t.Logf("disk io queue length: %v", diskIo.QueueLength)
	}
}

func TestGetDiskQueueLengthContinuously(t *testing.T) {
	collector, err := NewDiskQueueLengthCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	count := 0
	for {
		diskIoInfo, err := collector.GetValue()
		if err != nil {
			t.Fatalf("failed to collect metrics: %v", err)
		}
		diskIoInfos, ok := diskIoInfo.([]*DiskQueueLengthMetrics)
		if !ok {
			t.Fatalf("disk io info is not a []*DiskQueueLengthMetrics")
		}
		for _, diskIo := range diskIoInfos {
			t.Logf("disk io: %v", diskIo)
			t.Logf("disk io queue length: %v", diskIo.QueueLength)
		}
		count++
		if count >= 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
