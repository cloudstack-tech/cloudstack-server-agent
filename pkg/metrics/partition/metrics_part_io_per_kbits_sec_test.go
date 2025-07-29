package metrics

import (
	"testing"
	"time"
)

func TestGetPartitionIoPerKbitsSec(t *testing.T) {
	collector, err := NewPartitionIoPerKbitsSecCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	partIoInfo, err := collector.GetValue()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}
	partIoInfos, ok := partIoInfo.([]*PartitionIoPerKbitsSecMetrics)
	if !ok {
		t.Fatalf("part io info is not a []*PartitionIoPerKbitsSecMetrics")
	}
	for _, partIo := range partIoInfos {
		t.Logf("part io: %v", partIo)
		t.Logf("part io read : %v KB/s", partIo.KbitsRead/8)
		t.Logf("part io write: %v KB/s", partIo.KbitsWrite/8)
	}
}

func TestGetPartitionIoPerKbitsSecContinuously(t *testing.T) {
	collector, err := NewPartitionIoPerKbitsSecCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	count := 0
	for {
		partIoInfo, err := collector.GetValue()
		if err != nil {
			t.Fatalf("failed to collect metrics: %v", err)
		}
		partIoInfos, ok := partIoInfo.([]*PartitionIoPerKbitsSecMetrics)
		if !ok {
			t.Fatalf("part io info is not a []*PartitionIoPerKbitsSecMetrics")
		}
		for _, partIo := range partIoInfos {
			t.Logf("part io: %v", partIo)
			t.Logf("part io read : %v KB/s", partIo.KbitsRead/8)
			t.Logf("part io write: %v KB/s", partIo.KbitsWrite/8)
		}
		count++
		if count >= 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
