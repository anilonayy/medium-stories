package semaphore

import (
	"context"
	"crypto/sha256"
	"fmt"
	"sync"

	"golang.org/x/sync/semaphore"
)

// ProcessWithSemaphore limits concurrency using a weighted semaphore.
func ProcessWithSemaphore(ctx context.Context, items []string, maxConcurrency int64) []string {
	sem := semaphore.NewWeighted(maxConcurrency)
	var (
		mu      sync.Mutex
		results []string
		wg      sync.WaitGroup
	)

	wg.Add(len(items))
	for _, item := range items {
		if err := sem.Acquire(ctx, 1); err != nil {
			break
		}
		go func() {
			defer wg.Done()
			defer sem.Release(1)

			hash := sha256.Sum256([]byte(item))
			result := fmt.Sprintf("%x", hash[:8])

			mu.Lock()
			results = append(results, result)
			mu.Unlock()
		}()
	}

	wg.Wait()
	return results
}

// ProcessWithChannelSemaphore limits concurrency using a buffered channel.
func ProcessWithChannelSemaphore(items []string, maxConcurrency int) []string {
	sem := make(chan struct{}, maxConcurrency)
	var (
		mu      sync.Mutex
		results []string
		wg      sync.WaitGroup
	)

	wg.Add(len(items))
	for _, item := range items {
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()

			hash := sha256.Sum256([]byte(item))
			result := fmt.Sprintf("%x", hash[:8])

			mu.Lock()
			results = append(results, result)
			mu.Unlock()
		}()
	}

	wg.Wait()
	return results
}

// ProcessUnbounded spawns all goroutines at once — no concurrency limit.
func ProcessUnbounded(items []string) []string {
	var (
		mu      sync.Mutex
		results []string
		wg      sync.WaitGroup
	)

	wg.Add(len(items))
	for _, item := range items {
		go func() {
			defer wg.Done()

			hash := sha256.Sum256([]byte(item))
			result := fmt.Sprintf("%x", hash[:8])

			mu.Lock()
			results = append(results, result)
			mu.Unlock()
		}()
	}

	wg.Wait()
	return results
}
