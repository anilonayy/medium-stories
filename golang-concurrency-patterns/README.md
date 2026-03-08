# Concurrency Patterns Beyond Goroutines and Channels

Code examples and benchmarks for the Medium article: **"Concurrency Patterns Beyond Goroutines and Channels in Go"**

## Patterns Covered

| Pattern | Directory | What It Solves |
|---------|-----------|----------------|
| errgroup | `patterns/errgroup/` | Parallel tasks with structured error handling |
| singleflight | `patterns/singleflight/` | Cache stampede / thundering herd prevention |
| sync.Pool | `patterns/syncpool/` | Reduce GC pressure via object reuse |
| Worker Pool | `patterns/workerpool/` | Bounded concurrency for CPU/IO-bound work |
| Fan-Out/Fan-In | `patterns/fanout/` | Distribute work and merge results |
| Semaphore | `patterns/semaphore/` | Weighted concurrency control |
| sync.Map vs Mutex | `patterns/syncmap/` | Concurrent map access strategies |
| Rate Limiting | `patterns/ratelimit/` | Token bucket rate limiter |
| Context Patterns | `patterns/context/` | Timeout, cancellation, cascading cancel |

## Requirements

- Go 1.23+

## Run All Benchmarks

```bash
go test -bench=. -benchmem ./...
```

## Run a Specific Pattern

```bash
go test -bench=. -benchmem ./patterns/syncpool/
go test -bench=. -benchmem ./patterns/workerpool/
```

## Run with Race Detector

```bash
go test -race ./...
```

## Project Structure

```
golang-concurrency-patterns/
├── article.md                    # The full Medium article
├── patterns/
│   ├── errgroup/
│   │   ├── errgroup.go           # Sequential vs errgroup comparison
│   │   └── errgroup_test.go      # Benchmarks
│   ├── singleflight/
│   │   ├── singleflight.go       # Thundering herd demo
│   │   └── singleflight_test.go  # Benchmarks + dedup verification
│   ├── syncpool/
│   │   ├── syncpool.go           # Pool vs fresh allocation
│   │   └── syncpool_test.go      # Serial + parallel benchmarks
│   ├── workerpool/
│   │   ├── workerpool.go         # Unbounded vs fixed pool
│   │   └── workerpool_test.go    # 100 and 10K task benchmarks
│   ├── fanout/
│   │   ├── fanout.go             # Fan-out/fan-in vs sequential
│   │   └── fanout_test.go        # Benchmarks
│   ├── semaphore/
│   │   ├── semaphore.go          # Weighted vs channel vs unbounded
│   │   └── semaphore_test.go     # Benchmarks
│   ├── syncmap/
│   │   ├── syncmap.go            # sync.Map vs map+RWMutex
│   │   └── syncmap_test.go       # Read-heavy, mixed, write-heavy
│   ├── ratelimit/
│   │   ├── ratelimit.go          # Token bucket vs ticker vs none
│   │   └── ratelimit_test.go     # Benchmarks + accuracy tests
│   └── context/
│       ├── context.go            # Timeout, cascading cancellation
│       └── context_test.go       # Behavior tests + benchmarks
├── go.mod
└── go.sum
```
