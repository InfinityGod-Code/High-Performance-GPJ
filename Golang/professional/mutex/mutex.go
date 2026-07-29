package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// ====================================================================
// BASIC MUTEX — prevents multiple goroutines from accessing shared data
// ====================================================================
func mutexCounter() {
	var mu sync.Mutex
	count := 0
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++ // critical section
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Mutex counter (expected 100):", count)
}

// ====================================================================
// RWMUTEX — multiple readers OR a single writer
// ====================================================================
func rwMutexDemo() {
	var rw sync.RWMutex
	data := 0
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		rw.Lock()
		data = 42
		rw.Unlock()
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			rw.RLock()
			fmt.Printf("Reader %d sees data=%d\n", id, data)
			rw.RUnlock()
		}(i)
	}

	wg.Wait()
}

// ====================================================================
// COMPARISON: Race vs Mutex vs Atomic
// Run with: go run -race .
// ====================================================================
func raceVsMutexVsAtomic() {
	const ops = 10000
	var wg sync.WaitGroup

	racyCount := 0
	for i := 0; i < ops; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			racyCount++ // DATA RACE
		}()
	}
	wg.Wait()
	fmt.Println("Race count (likely wrong):", racyCount)

	mutexCount := 0
	var mu sync.Mutex
	for i := 0; i < ops; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			mutexCount++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Mutex count (expected", ops, "):", mutexCount)

	var atomicCount int64
	for i := 0; i < ops; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&atomicCount, 1)
		}()
	}
	wg.Wait()
	fmt.Println("Atomic count (expected", ops, "):", atomicCount)
}

func main() {
	fmt.Println("=== Basic Mutex Counter ===")
	mutexCounter()

	fmt.Println("\n=== RWMutex Demo ===")
	rwMutexDemo()

	fmt.Println("\n=== Race vs Mutex vs Atomic ===")
	raceVsMutexVsAtomic()
}
