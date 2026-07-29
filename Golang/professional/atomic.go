package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

// ====================================================================
// ATOMIC ADD — safely increments a counter across multiple goroutines
// ====================================================================
func atomicAdd() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Atomic counter (expected 100):", counter)
}

// ====================================================================
// ATOMIC LOAD & STORE — safely read / write a shared value
// ====================================================================
func atomicLoadStore() {
	var shared int64
	var wg sync.WaitGroup

	// Writer goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		atomic.StoreInt64(&shared, 42)
	}()

	// Reader goroutine
	var readValue int64
	wg.Add(1)
	go func() {
		defer wg.Done()
		readValue = atomic.LoadInt64(&shared)
	}()

	wg.Wait()
	fmt.Println("Loaded atomic value:", readValue)
}

// ====================================================================
// ATOMIC COMPARE AND SWAP (CAS) — update only if value hasn't changed
// ====================================================================
func atomicCAS() {
	var val int64 = 10

	// Attempt to swap 10 -> 20 only if current value is 10
	swapped := atomic.CompareAndSwapInt64(&val, 10, 20)
	fmt.Printf("CAS (10->20): swapped=%v, val=%d\n", swapped, val)

	// This one will fail because val is now 20, not 10
	swapped = atomic.CompareAndSwapInt64(&val, 10, 30)
	fmt.Printf("CAS (10->30): swapped=%v, val=%d\n", swapped, val)
}

// ====================================================================
// ATOMIC SWAP — atomically exchange a value and return the old one
// ====================================================================
func atomicSwap() {
	var val int64 = 100
	old := atomic.SwapInt64(&val, 200)
	fmt.Printf("Swap: old=%d, new=%d\n", old, val)
}
