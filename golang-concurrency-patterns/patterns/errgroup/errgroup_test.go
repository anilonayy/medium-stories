package errgroup

import (
	"context"
	"testing"
)

func BenchmarkFetchDashboardNaive(b *testing.B) {
	ctx := context.Background()
	for b.Loop() {
		_, err := FetchDashboardNaive(ctx)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFetchDashboardErrgroup(b *testing.B) {
	ctx := context.Background()
	for b.Loop() {
		_, err := FetchDashboardErrgroup(ctx)
		if err != nil {
			b.Fatal(err)
		}
	}
}
