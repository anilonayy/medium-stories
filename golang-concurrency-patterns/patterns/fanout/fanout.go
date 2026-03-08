package fanout

import (
	"context"
	"crypto/sha256"
	"fmt"
	"sync"
)

// FanOutFanIn distributes work across numWorkers goroutines and merges results.
func FanOutFanIn(ctx context.Context, inputs []string, numWorkers int) []string {
	inputCh := make(chan string, len(inputs))
	for _, in := range inputs {
		inputCh <- in
	}
	close(inputCh)

	// Fan-out: multiple workers read from the same channel
	resultChannels := make([]<-chan string, numWorkers)
	for i := range numWorkers {
		resultChannels[i] = worker(ctx, inputCh)
	}

	// Fan-in: merge all worker outputs into one channel
	merged := merge(resultChannels...)

	var results []string
	for r := range merged {
		results = append(results, r)
	}
	return results
}

func worker(ctx context.Context, input <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for data := range input {
			select {
			case <-ctx.Done():
				return
			default:
				hash := sha256.Sum256([]byte(data))
				out <- fmt.Sprintf("%x", hash[:8])
			}
		}
	}()
	return out
}

func merge(channels ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	merged := make(chan string)

	wg.Add(len(channels))
	for _, ch := range channels {
		go func() {
			defer wg.Done()
			for val := range ch {
				merged <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

// Sequential processes all inputs in a single goroutine.
func Sequential(inputs []string) []string {
	results := make([]string, len(inputs))
	for i, data := range inputs {
		hash := sha256.Sum256([]byte(data))
		results[i] = fmt.Sprintf("%x", hash[:8])
	}
	return results
}
