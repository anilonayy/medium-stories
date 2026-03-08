package ratelimit

import (
	"context"
	"testing"
	"time"
)

func BenchmarkProcessWithRateLimiter(b *testing.B) {
	ctx := context.Background()
	// High rate to measure overhead, not wait time
	for b.Loop() {
		ProcessWithRateLimiter(ctx, 100, 100_000, 100)
	}
}

func BenchmarkProcessWithoutRateLimiter(b *testing.B) {
	for b.Loop() {
		ProcessWithoutRateLimiter(100)
	}
}

func TestRateLimiterAccuracy(t *testing.T) {
	ctx := context.Background()
	// Allow 100 requests/sec with burst of 1
	processed, duration := ProcessWithRateLimiter(ctx, 10, 100, 1)

	expectedMin := 80 * time.Millisecond // ~90ms for 9 waits at 100/s
	if duration < expectedMin {
		t.Errorf("rate limiter too fast: processed %d in %v (expected >= %v)", processed, duration, expectedMin)
	}
	t.Logf("processed %d items in %v (target: 100 req/s)", processed, duration)
}

func TestManualTickerAccuracy(t *testing.T) {
	processed, duration := ProcessWithManualTicker(10, 10*time.Millisecond)

	expectedMin := 90 * time.Millisecond
	if duration < expectedMin {
		t.Errorf("ticker too fast: processed %d in %v (expected >= %v)", processed, duration, expectedMin)
	}
	t.Logf("processed %d items in %v (target: 10ms interval)", processed, duration)
}
