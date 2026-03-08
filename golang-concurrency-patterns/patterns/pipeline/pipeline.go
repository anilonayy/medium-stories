package pipeline

import (
	"context"
	"strings"
)

// Stage 1: Generate — produces data from a slice.
func Generate(ctx context.Context, data []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, d := range data {
			select {
			case out <- d:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// Stage 2: Transform — processes each item (e.g., uppercase + trim).
func Transform(ctx context.Context, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range in {
			select {
			case <-ctx.Done():
				return
			default:
				out <- strings.ToUpper(strings.TrimSpace(s))
			}
		}
	}()
	return out
}

// Stage 3: Filter — keeps only items matching a condition.
func Filter(ctx context.Context, in <-chan string, minLen int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range in {
			select {
			case <-ctx.Done():
				return
			default:
				if len(s) >= minLen {
					out <- s
				}
			}
		}
	}()
	return out
}

// RunPipeline connects the three stages into a pipeline.
func RunPipeline(ctx context.Context, data []string, minLen int) []string {
	// Stage 1 -> Stage 2 -> Stage 3
	stage1 := Generate(ctx, data)
	stage2 := Transform(ctx, stage1)
	stage3 := Filter(ctx, stage2, minLen)

	var results []string
	for r := range stage3 {
		results = append(results, r)
	}
	return results
}

// RunSequential does the same work without concurrency.
func RunSequential(data []string, minLen int) []string {
	var results []string
	for _, d := range data {
		transformed := strings.ToUpper(strings.TrimSpace(d))
		if len(transformed) >= minLen {
			results = append(results, transformed)
		}
	}
	return results
}
