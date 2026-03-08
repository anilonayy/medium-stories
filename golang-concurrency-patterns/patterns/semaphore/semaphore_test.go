package semaphore

import (
	"context"
	"fmt"
	"testing"
)

func generateItems(n int) []string {
	items := make([]string, n)
	for i := range n {
		items[i] = fmt.Sprintf("item-data-%d-for-processing", i)
	}
	return items
}

func BenchmarkProcessUnbounded_1000(b *testing.B) {
	items := generateItems(1_000)
	for b.Loop() {
		_ = ProcessUnbounded(items)
	}
}

func BenchmarkProcessWithSemaphore_1000(b *testing.B) {
	items := generateItems(1_000)
	ctx := context.Background()
	for b.Loop() {
		_ = ProcessWithSemaphore(ctx, items, 10)
	}
}

func BenchmarkProcessWithChannelSemaphore_1000(b *testing.B) {
	items := generateItems(1_000)
	for b.Loop() {
		_ = ProcessWithChannelSemaphore(items, 10)
	}
}

func BenchmarkProcessUnbounded_10000(b *testing.B) {
	items := generateItems(10_000)
	for b.Loop() {
		_ = ProcessUnbounded(items)
	}
}

func BenchmarkProcessWithSemaphore_10000(b *testing.B) {
	items := generateItems(10_000)
	ctx := context.Background()
	for b.Loop() {
		_ = ProcessWithSemaphore(ctx, items, 10)
	}
}

func BenchmarkProcessWithChannelSemaphore_10000(b *testing.B) {
	items := generateItems(10_000)
	for b.Loop() {
		_ = ProcessWithChannelSemaphore(items, 10)
	}
}
