package ratelimit

import (
	"context"
	"sync/atomic"
	"time"

	"golang.org/x/time/rate"
)

// ProcessWithRateLimiter processes items respecting a rate limit.
func ProcessWithRateLimiter(ctx context.Context, count int, rps float64, burst int) (int64, time.Duration) {
	limiter := rate.NewLimiter(rate.Limit(rps), burst)
	var processed atomic.Int64
	start := time.Now()

	for range count {
		if err := limiter.Wait(ctx); err != nil {
			break
		}
		processed.Add(1)
	}

	return processed.Load(), time.Since(start)
}

// ProcessWithoutRateLimiter processes items as fast as possible.
func ProcessWithoutRateLimiter(count int) (int64, time.Duration) {
	var processed atomic.Int64
	start := time.Now()

	for range count {
		processed.Add(1)
	}

	return processed.Load(), time.Since(start)
}

// ProcessWithManualTicker uses a time.Ticker for rate limiting — the naive approach.
func ProcessWithManualTicker(count int, interval time.Duration) (int64, time.Duration) {
	var processed atomic.Int64
	start := time.Now()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range count {
		<-ticker.C
		processed.Add(1)
	}

	return processed.Load(), time.Since(start)
}
