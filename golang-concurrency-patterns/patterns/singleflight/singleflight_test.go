package singleflight

import (
	"testing"
)

func BenchmarkFetchWithoutSingleflight(b *testing.B) {
	for b.Loop() {
		_, err := FetchWithoutSingleflight("user:123", 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFetchWithSingleflight(b *testing.B) {
	for b.Loop() {
		_, err := FetchWithSingleflight("user:123", 100)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestSingleflightReducesCalls(t *testing.T) {
	_, err := FetchWithSingleflight("user:456", 50)
	if err != nil {
		t.Fatal(err)
	}
	calls := GetDBCallCount()
	if calls > 5 {
		t.Errorf("expected at most 5 DB calls, got %d", calls)
	}
	t.Logf("singleflight reduced 50 concurrent requests to %d DB call(s)", calls)
}
