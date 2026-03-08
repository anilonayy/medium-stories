package fanout

import (
	"context"
	"fmt"
	"runtime"
	"testing"
)

func generateInputs(n int) []string {
	inputs := make([]string, n)
	for i := range n {
		inputs[i] = fmt.Sprintf("data-payload-%d-for-hashing", i)
	}
	return inputs
}

func BenchmarkSequential_1000(b *testing.B) {
	inputs := generateInputs(1_000)
	for b.Loop() {
		_ = Sequential(inputs)
	}
}

func BenchmarkFanOutFanIn_1000(b *testing.B) {
	inputs := generateInputs(1_000)
	ctx := context.Background()
	numWorkers := runtime.GOMAXPROCS(0)
	for b.Loop() {
		_ = FanOutFanIn(ctx, inputs, numWorkers)
	}
}

func BenchmarkSequential_10000(b *testing.B) {
	inputs := generateInputs(10_000)
	for b.Loop() {
		_ = Sequential(inputs)
	}
}

func BenchmarkFanOutFanIn_10000(b *testing.B) {
	inputs := generateInputs(10_000)
	ctx := context.Background()
	numWorkers := runtime.GOMAXPROCS(0)
	for b.Loop() {
		_ = FanOutFanIn(ctx, inputs, numWorkers)
	}
}
