package ctx

import (
	"testing"
	"time"
)

func TestWorkWithTimeout_CompletesEarly(t *testing.T) {
	result := WorkWithTimeout(500*time.Millisecond, 10)
	if result.Cancelled {
		t.Error("expected work to complete without cancellation")
	}
	t.Logf("completed %d items in %v", result.Completed, result.Duration)
}

func TestWorkWithTimeout_TimesOut(t *testing.T) {
	result := WorkWithTimeout(20*time.Millisecond, 1000)
	if !result.Cancelled {
		t.Error("expected work to be cancelled due to timeout")
	}
	t.Logf("completed %d/1000 items before timeout in %v", result.Completed, result.Duration)
}

func TestCascadingCancel(t *testing.T) {
	parent, child := WorkWithCascadingCancel(50*time.Millisecond, 200*time.Millisecond)

	t.Logf("parent completed %d items", parent.Completed)
	t.Logf("child completed %d items (child timeout was 200ms, but parent was 50ms)", child.Completed)

	// Child should have been cancelled by parent, not by its own timeout
	if child.Completed > parent.Completed+5 {
		t.Error("child should have been cancelled when parent was cancelled")
	}
}

func BenchmarkWorkWithTimeout(b *testing.B) {
	for b.Loop() {
		_ = WorkWithTimeout(100*time.Millisecond, 50)
	}
}

func BenchmarkWorkIgnoringContext(b *testing.B) {
	for b.Loop() {
		_ = WorkIgnoringContext(100*time.Millisecond, 50)
	}
}
