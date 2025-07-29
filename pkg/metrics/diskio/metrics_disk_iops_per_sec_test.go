package metrics

import (
	"testing"
	"time"
)

func TestGetDiskIoPerIopsSec(t *testing.T) {
	collector, err := NewDiskIoPerIopsSecCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	diskIoInfo, err := collector.GetValue()
	if err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}
	diskIoInfos, ok := diskIoInfo.([]*DiskIoPerIopsSecMetrics)
	if !ok {
		t.Fatalf("disk io info is not a []*DiskIoPerIopsSecMetrics")
	}
	for _, diskIo := range diskIoInfos {
		t.Logf("disk io: %v", diskIo)
		t.Logf("disk io read : %v iops", diskIo.IopsRead)
		t.Logf("disk io write: %v iops", diskIo.IopsWrite)
	}
}

func TestGetDiskIoPerIopsSecContinuously(t *testing.T) {
	collector, err := NewDiskIoPerIopsSecCollector()
	if err != nil {
		t.Fatalf("failed to create collector: %v", err)
	}
	count := 0
	for {
		diskIoInfo, err := collector.GetValue()
		if err != nil {
			t.Fatalf("failed to collect metrics: %v", err)
		}
		diskIoInfos, ok := diskIoInfo.([]*DiskIoPerIopsSecMetrics)
		if !ok {
			t.Fatalf("disk io info is not a []*DiskIoPerIopsSecMetrics")
		}
		for _, diskIo := range diskIoInfos {
			t.Logf("disk io: %v", diskIo)
			t.Logf("disk io read : %v iops", diskIo.IopsRead)
			t.Logf("disk io write: %v iops", diskIo.IopsWrite)
		}
		count++
		if count >= 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
