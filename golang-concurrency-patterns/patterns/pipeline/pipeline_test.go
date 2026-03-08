package pipeline

import (
	"context"
	"fmt"
	"testing"
)

func generateData(n int) []string {
	data := make([]string, n)
	for i := range n {
		data[i] = fmt.Sprintf("  item number %d with some padding  ", i)
	}
	return data
}

func BenchmarkRunSequential_1000(b *testing.B) {
	data := generateData(1_000)
	for b.Loop() {
		_ = RunSequential(data, 10)
	}
}

func BenchmarkRunPipeline_1000(b *testing.B) {
	data := generateData(1_000)
	ctx := context.Background()
	for b.Loop() {
		_ = RunPipeline(ctx, data, 10)
	}
}

func BenchmarkRunSequential_50000(b *testing.B) {
	data := generateData(50_000)
	for b.Loop() {
		_ = RunSequential(data, 10)
	}
}

func BenchmarkRunPipeline_50000(b *testing.B) {
	data := generateData(50_000)
	ctx := context.Background()
	for b.Loop() {
		_ = RunPipeline(ctx, data, 10)
	}
}
