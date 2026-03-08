package workerpool

import (
	"context"
	"crypto/sha256"
	"fmt"
	"sync"
)

// Task represents a unit of work.
type Task struct {
	ID   int
	Data []byte
}

// Result represents a processed result.
type Result struct {
	TaskID int
	Hash   string
}

// ProcessUnbounded spawns a goroutine per task — no concurrency control.
func ProcessUnbounded(tasks []Task) []Result {
	var (
		mu      sync.Mutex
		results []Result
		wg      sync.WaitGroup
	)

	wg.Add(len(tasks))
	for _, t := range tasks {
		go func() {
			defer wg.Done()
			hash := sha256.Sum256(t.Data)
			mu.Lock()
			results = append(results, Result{
				TaskID: t.ID,
				Hash:   fmt.Sprintf("%x", hash),
			})
			mu.Unlock()
		}()
	}

	wg.Wait()
	return results
}

// ProcessWorkerPool uses a fixed number of workers to process tasks.
func ProcessWorkerPool(ctx context.Context, tasks []Task, numWorkers int) []Result {
	taskCh := make(chan Task, len(tasks))
	resultCh := make(chan Result, len(tasks))

	// Start workers
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for range numWorkers {
		go func() {
			defer wg.Done()
			for t := range taskCh {
				select {
				case <-ctx.Done():
					return
				default:
					hash := sha256.Sum256(t.Data)
					resultCh <- Result{
						TaskID: t.ID,
						Hash:   fmt.Sprintf("%x", hash),
					}
				}
			}
		}()
	}

	// Send tasks
	for _, t := range tasks {
		taskCh <- t
	}
	close(taskCh)

	// Wait for workers and close results
	wg.Wait()
	close(resultCh)

	var results []Result
	for r := range resultCh {
		results = append(results, r)
	}
	return results
}
