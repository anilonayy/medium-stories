package syncmap

import "testing"

// 95% reads, 5% writes — sync.Map's sweet spot
func BenchmarkReadHeavyWithMutex_95Read(b *testing.B) {
	for b.Loop() {
		ReadHeavyWithMutex(95, 5, 100)
	}
}

func BenchmarkReadHeavyWithSyncMap_95Read(b *testing.B) {
	for b.Loop() {
		ReadHeavyWithSyncMap(95, 5, 100)
	}
}

// 50% reads, 50% writes — mixed workload
func BenchmarkMixedWithMutex_50Read(b *testing.B) {
	for b.Loop() {
		ReadHeavyWithMutex(50, 50, 100)
	}
}

func BenchmarkMixedWithSyncMap_50Read(b *testing.B) {
	for b.Loop() {
		ReadHeavyWithSyncMap(50, 50, 100)
	}
}

// 100% writes — worst case for sync.Map
func BenchmarkWriteHeavyWithMutex(b *testing.B) {
	for b.Loop() {
		ReadHeavyWithMutex(0, 100, 100)
	}
}

func BenchmarkWriteHeavyWithSyncMap(b *testing.B) {
	for b.Loop() {
		ReadHeavyWithSyncMap(0, 100, 100)
	}
}
