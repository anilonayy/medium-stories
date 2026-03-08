package ctx

import (
	"context"
	"sync/atomic"
	"time"
)

// WorkResult represents the outcome of a context-controlled operation.
type WorkResult struct {
	Completed int64
	Cancelled bool
	Duration  time.Duration
}

// WorkWithTimeout performs work that respects a timeout.
func WorkWithTimeout(timeout time.Duration, totalItems int) WorkResult {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var completed atomic.Int64
	start := time.Now()

	for range totalItems {
		select {
		case <-ctx.Done():
			return WorkResult{
				Completed: completed.Load(),
				Cancelled: true,
				Duration:  time.Since(start),
			}
		default:
			// Simulate work
			time.Sleep(1 * time.Millisecond)
			completed.Add(1)
		}
	}

	return WorkResult{
		Completed: completed.Load(),
		Cancelled: false,
		Duration:  time.Since(start),
	}
}

// WorkWithCascadingCancel demonstrates parent-child context cancellation.
func WorkWithCascadingCancel(parentTimeout, childTimeout time.Duration) (parent WorkResult, child WorkResult) {
	parentCtx, parentCancel := context.WithTimeout(context.Background(), parentTimeout)
	defer parentCancel()

	childCtx, childCancel := context.WithTimeout(parentCtx, childTimeout)
	defer childCancel()

	start := time.Now()

	// Parent work
	var parentCompleted atomic.Int64
	go func() {
		for {
			select {
			case <-parentCtx.Done():
				return
			default:
				time.Sleep(1 * time.Millisecond)
				parentCompleted.Add(1)
			}
		}
	}()

	// Child work — constrained by BOTH parent and child timeouts
	var childCompleted atomic.Int64
	childDone := make(chan struct{})
	go func() {
		defer close(childDone)
		for {
			select {
			case <-childCtx.Done():
				return
			default:
				time.Sleep(1 * time.Millisecond)
				childCompleted.Add(1)
			}
		}
	}()

	// Wait for parent to finish (which also cancels child)
	<-parentCtx.Done()
	<-childDone

	elapsed := time.Since(start)
	parent = WorkResult{Completed: parentCompleted.Load(), Cancelled: true, Duration: elapsed}
	child = WorkResult{Completed: childCompleted.Load(), Cancelled: true, Duration: elapsed}
	return
}

// WorkIgnoringContext does work without checking context — the anti-pattern.
func WorkIgnoringContext(timeout time.Duration, totalItems int) WorkResult {
	start := time.Now()
	var completed int64

	for range totalItems {
		time.Sleep(1 * time.Millisecond)
		completed++
	}

	return WorkResult{
		Completed: completed,
		Cancelled: false,
		Duration:  time.Since(start),
	}
}
