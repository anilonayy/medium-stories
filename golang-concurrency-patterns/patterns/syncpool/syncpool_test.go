package syncpool

import (
	"bytes"
	"testing"
)

var testData = bytes.Repeat([]byte("hello world, this is a concurrency benchmark test "), 100)

func BenchmarkProcessWithPool(b *testing.B) {
	for b.Loop() {
		_ = ProcessWithPool(testData)
	}
}

func BenchmarkProcessWithoutPool(b *testing.B) {
	for b.Loop() {
		_ = ProcessWithoutPool(testData)
	}
}

func BenchmarkProcessWithPoolParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = ProcessWithPool(testData)
		}
	})
}

func BenchmarkProcessWithoutPoolParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = ProcessWithoutPool(testData)
		}
	})
}
