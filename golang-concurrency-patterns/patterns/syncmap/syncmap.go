package syncmap

import (
	"fmt"
	"sync"
)

// ReadHeavyWithMutex uses a regular map with RWMutex for concurrent access.
func ReadHeavyWithMutex(readers, writers int, iterations int) {
	var mu sync.RWMutex
	m := make(map[string]string)

	// Pre-populate
	for i := range 1000 {
		m[fmt.Sprintf("key-%d", i)] = fmt.Sprintf("value-%d", i)
	}

	var wg sync.WaitGroup
	wg.Add(readers + writers)

	// Readers
	for range readers {
		go func() {
			defer wg.Done()
			for i := range iterations {
				key := fmt.Sprintf("key-%d", i%1000)
				mu.RLock()
				_ = m[key]
				mu.RUnlock()
			}
		}()
	}

	// Writers
	for range writers {
		go func() {
			defer wg.Done()
			for i := range iterations {
				key := fmt.Sprintf("key-%d", i%1000)
				mu.Lock()
				m[key] = fmt.Sprintf("updated-%d", i)
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
}

// ReadHeavyWithSyncMap uses sync.Map for concurrent access.
func ReadHeavyWithSyncMap(readers, writers int, iterations int) {
	var m sync.Map

	// Pre-populate
	for i := range 1000 {
		m.Store(fmt.Sprintf("key-%d", i), fmt.Sprintf("value-%d", i))
	}

	var wg sync.WaitGroup
	wg.Add(readers + writers)

	// Readers
	for range readers {
		go func() {
			defer wg.Done()
			for i := range iterations {
				key := fmt.Sprintf("key-%d", i%1000)
				_, _ = m.Load(key)
			}
		}()
	}

	// Writers
	for range writers {
		go func() {
			defer wg.Done()
			for i := range iterations {
				key := fmt.Sprintf("key-%d", i%1000)
				m.Store(key, fmt.Sprintf("updated-%d", i))
			}
		}()
	}

	wg.Wait()
}
