# medium-stories

Companion code for my Medium articles. Each article gets its own folder with runnable examples and benchmarks.

## Articles

| Folder | Article | Contents |
|--------|---------|----------|
| [`golang-concurrency-patterns/`](golang-concurrency-patterns/) | Concurrency Patterns Beyond Goroutines and Channels in Go | Go examples and benchmarks for errgroup, singleflight, sync.Pool, worker pool, fan-out/fan-in, semaphore, sync.Map vs mutex, rate limiting, context and pipeline |

## Running the Go examples

Requires Go (see the `go` directive in `golang-concurrency-patterns/go.mod`).

```bash
cd golang-concurrency-patterns
go test -bench=. -benchmem ./...   # all benchmarks
go test -race ./...                # with the race detector
```

See [`golang-concurrency-patterns/README.md`](golang-concurrency-patterns/README.md) for the per-pattern breakdown.
