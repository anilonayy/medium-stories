package singleflight

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

// simulateDBQuery simulates an expensive database call.
func simulateDBQuery(key string) (string, error) {
	time.Sleep(time.Duration(5+rand.Intn(10)) * time.Millisecond)
	return fmt.Sprintf("data-for-%s", key), nil
}

var dbCallCount atomic.Int64

// FetchWithoutSingleflight simulates concurrent cache misses hitting the DB directly.
func FetchWithoutSingleflight(key string, concurrency int) ([]string, error) {
	dbCallCount.Store(0)

	var (
		mu      sync.Mutex
		results []string
		wg      sync.WaitGroup
	)

	wg.Add(concurrency)
	for range concurrency {
		go func() {
			defer wg.Done()
			dbCallCount.Add(1)
			val, err := simulateDBQuery(key)
			if err != nil {
				return
			}
			mu.Lock()
			results = append(results, val)
			mu.Unlock()
		}()
	}
	wg.Wait()

	return results, nil
}

// FetchWithSingleflight deduplicates concurrent calls for the same key.
func FetchWithSingleflight(key string, concurrency int) ([]string, error) {
	dbCallCount.Store(0)

	var g singleflight.Group
	var (
		mu      sync.Mutex
		results []string
		wg      sync.WaitGroup
	)

	wg.Add(concurrency)
	for range concurrency {
		go func() {
			defer wg.Done()
			val, _, _ := g.Do(key, func() (interface{}, error) {
				dbCallCount.Add(1)
				return simulateDBQuery(key)
			})
			mu.Lock()
			results = append(results, val.(string))
			mu.Unlock()
		}()
	}
	wg.Wait()

	return results, nil
}

// GetDBCallCount returns the number of actual DB calls made.
func GetDBCallCount() int64 {
	return dbCallCount.Load()
}
