package metrics

import (
	"testing"
	"time"
)

func TestGetDiskIoPerKbitsSec(t *testing.T) {
	collector, err := NewDiskIoPerKbitsSecCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	diskIoInfo, err := collector.GetValue()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}
	diskIoInfos, ok := diskIoInfo.([]*DiskIoPerKbitsSecMetrics)
	if !ok {
		t.Fatalf("disk io info is not a []*DiskIoPerKbitsSecMetrics")
	}
	for _, diskIo := range diskIoInfos {
		t.Logf("disk io: %v", diskIo)
		t.Logf("disk io read : %v KB/s", diskIo.KbitsRead/8)
		t.Logf("disk io write: %v KB/s", diskIo.KbitsWrite/8)
	}
}

func TestGetDiskIoPerKbitsSecContinuously(t *testing.T) {
	collector, err := NewDiskIoPerKbitsSecCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	count := 0
	for {
		diskIoInfo, err := collector.GetValue()
		if err != nil {
			t.Fatalf("failed to collect metrics: %v", err)
		}
		diskIoInfos, ok := diskIoInfo.([]*DiskIoPerKbitsSecMetrics)
		if !ok {
			t.Fatalf("disk io info is not a []*DiskIoPerKbitsSecMetrics")
		}
		for _, diskIo := range diskIoInfos {
			t.Logf("disk io: %v", diskIo)
			t.Logf("disk io read : %v KB/s", diskIo.KbitsRead/8)
			t.Logf("disk io write: %v KB/s", diskIo.KbitsWrite/8)
		}
		count++
		if count >= 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
