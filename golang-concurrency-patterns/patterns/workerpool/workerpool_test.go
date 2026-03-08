package workerpool

import (
	"context"
	"fmt"
	"runtime"
	"testing"
)

func generateTasks(n int) []Task {
	tasks := make([]Task, n)
	for i := range n {
		tasks[i] = Task{
			ID:   i,
			Data: []byte(fmt.Sprintf("payload-%d-with-some-extra-data-to-hash", i)),
		}
	}
	return tasks
}

func BenchmarkProcessUnbounded_100(b *testing.B) {
	tasks := generateTasks(100)
	for b.Loop() {
		_ = ProcessUnbounded(tasks)
	}
}

func BenchmarkProcessUnbounded_10000(b *testing.B) {
	tasks := generateTasks(10_000)
	for b.Loop() {
		_ = ProcessUnbounded(tasks)
	}
}

func BenchmarkProcessWorkerPool_100(b *testing.B) {
	tasks := generateTasks(100)
	ctx := context.Background()
	numWorkers := runtime.GOMAXPROCS(0)
	for b.Loop() {
		_ = ProcessWorkerPool(ctx, tasks, numWorkers)
	}
}

func BenchmarkProcessWorkerPool_10000(b *testing.B) {
	tasks := generateTasks(10_000)
	ctx := context.Background()
	numWorkers := runtime.GOMAXPROCS(0)
	for b.Loop() {
		_ = ProcessWorkerPool(ctx, tasks, numWorkers)
	}
}
