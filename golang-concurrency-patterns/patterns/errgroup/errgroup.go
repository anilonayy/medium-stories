package errgroup

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

// UserProfile represents data from the user service.
type UserProfile struct {
	Name  string
	Email string
}

// Recommendation represents data from the recommendation service.
type Recommendation struct {
	Items []string
}

// Notification represents data from the notification service.
type Notification struct {
	Count int
}

// DashboardData aggregates data from multiple services.
type DashboardData struct {
	Profile         UserProfile
	Recommendations Recommendation
	Notifications   Notification
}

// simulateServiceCall simulates an API call with variable latency.
func simulateServiceCall(ctx context.Context, name string) error {
	delay := time.Duration(50+rand.Intn(100)) * time.Millisecond
	select {
	case <-time.After(delay):
		return nil
	case <-ctx.Done():
		return fmt.Errorf("%s: %w", name, ctx.Err())
	}
}

// FetchDashboardNaive fetches dashboard data sequentially — the slow way.
func FetchDashboardNaive(ctx context.Context) (*DashboardData, error) {
	var data DashboardData

	if err := simulateServiceCall(ctx, "profile"); err != nil {
		return nil, err
	}
	data.Profile = UserProfile{Name: "Jane", Email: "jane@example.com"}

	if err := simulateServiceCall(ctx, "recommendations"); err != nil {
		return nil, err
	}
	data.Recommendations = Recommendation{Items: []string{"item1", "item2"}}

	if err := simulateServiceCall(ctx, "notifications"); err != nil {
		return nil, err
	}
	data.Notifications = Notification{Count: 5}

	return &data, nil
}

// FetchDashboardErrgroup fetches dashboard data concurrently with errgroup.
func FetchDashboardErrgroup(ctx context.Context) (*DashboardData, error) {
	var data DashboardData
	var mu sync.Mutex

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		if err := simulateServiceCall(ctx, "profile"); err != nil {
			return err
		}
		mu.Lock()
		data.Profile = UserProfile{Name: "Jane", Email: "jane@example.com"}
		mu.Unlock()
		return nil
	})

	g.Go(func() error {
		if err := simulateServiceCall(ctx, "recommendations"); err != nil {
			return err
		}
		mu.Lock()
		data.Recommendations = Recommendation{Items: []string{"item1", "item2"}}
		mu.Unlock()
		return nil
	})

	g.Go(func() error {
		if err := simulateServiceCall(ctx, "notifications"); err != nil {
			return err
		}
		mu.Lock()
		data.Notifications = Notification{Count: 5}
		mu.Unlock()
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &data, nil
}
