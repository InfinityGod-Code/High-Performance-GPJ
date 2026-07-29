## Welcome
Welcome to the professional section of learning and mastering go before we start anything there is one famous chant that Go gives
                ** Don't communicate by sharing memory; share memory by communicating **
                ** Don't communicate by sharing memory; share memory by communicating **

---

## Table of Contents
1. [Goroutines & WaitGroups](#1-goroutines--waitgroups)
2. [Race Conditions](#2-race-conditions)
3. [Atomic Operations](#3-atomic-operations)
4. [Mutex & RWMutex](#4-mutex--rwmutex)
5. [Atomic vs Mutex — When to Use What](#5-atomic-vs-mutex--when-to-use-what)

---

## 1. Goroutines & WaitGroups

**Reference:** `waitgroup.go`

- The `main()` function itself runs as a goroutine.
- If `main` exits, **all** other goroutines are killed immediately — they don't get a chance to finish.
- `sync.WaitGroup` is a counter that blocks `Wait()` until the counter reaches zero.
  - `Add(n)` — increments the counter by `n` (call *before* launching the goroutine).
  - `Done()` — decrements the counter by 1 (call inside the goroutine, usually via `defer`).
  - `Wait()` — blocks until the counter is 0.

```go
wg := &sync.WaitGroup{}
wg.Add(1)
go func() {
    defer wg.Done()
    // do work
}()
wg.Wait()
```

> **Note:** WaitGroups must be passed **by reference** (pointer) to goroutines so they can call `Done()` on the same instance.

---

## 2. Race Conditions

**Reference:** `race_conditions.go`

A **data race** occurs when two or more goroutines access the same variable concurrently, and at least one of them is a write.

```go
func raceConditions(count *int) {
    c := *count     // read
    *count = c + 1  // write — unsafe without synchronisation
}
```

- Calling this function from multiple goroutines produces **unpredictable results**.
- The `main.go` example shows that even with 3 goroutines, the final count is often 0 because all goroutines read before any of them writes.

### Detecting Races

```bash
go run -race .
```

The `-race` flag instruments your binary to detect data races at runtime.

---

## 3. Atomic Operations

**Reference:** `atomic.go`

The `sync/atomic` package provides low-level atomic operations on primitive types. These are **lock-free** primitives that the CPU guarantees as a single, indivisible instruction.

### Operations covered:

| Function | Purpose |
|----------|---------|
| `atomic.AddT(ptr, delta)` | Atomically add `delta` to `*ptr` |
| `atomic.LoadT(ptr)` | Atomically read `*ptr` |
| `atomic.StoreT(ptr, val)` | Atomically write `val` to `*ptr` |
| `atomic.CompareAndSwapT(ptr, old, new)` | If `*ptr == old`, set it to `new` (CAS) |
| `atomic.SwapT(ptr, new)` | Atomically exchange and return the old value |

`T` can be `Int32`, `Int64`, `Uint32`, `Uint64`, `Pointer`, etc.

```go
atomic.AddInt64(&counter, 1)          // thread-safe increment
val := atomic.LoadInt64(&shared)      // safe read
atomic.StoreInt64(&shared, 42)        // safe write
swapped := atomic.CompareAndSwapInt64(&val, 10, 20) // CAS
old := atomic.SwapInt64(&val, 200)    // atomic swap
```

> **When to use atomic:** Simple counters, flags, and state indicators. Atomic operations are faster than mutexes because they avoid OS-level context switches.

---

## 4. Mutex & RWMutex

**Reference:** `mutex/mutex.go`

### sync.Mutex

A **mutual exclusion lock** ensures that only one goroutine can execute the **critical section** at a time.

```go
var mu sync.Mutex
mu.Lock()
count++ // critical section
mu.Unlock()
```

- `Lock()` — acquires the lock (blocks if another goroutine holds it).
- `Unlock()` — releases the lock.
- **Always** pair `Lock` with `Unlock` — prefer `defer mu.Unlock()` when the critical section has multiple exit paths.

### sync.RWMutex

A **read-write mutex** allows:
- **Multiple concurrent readers** (they don't block each other).
- **Exclusive access for a writer** (blocks all readers and other writers).

```go
var rw sync.RWMutex
rw.RLock()   // read lock — many goroutines can hold this simultaneously
// read data
rw.RUnlock()

rw.Lock()    // write lock — exclusive
// write data
rw.Unlock()
```

> **Use RWMutex** when reads are frequent and writes are infrequent. Otherwise, a plain Mutex is simpler and often just as fast.

---

## 5. Atomic vs Mutex — When to Use What

| Criteria | Atomic | Mutex |
|----------|--------|-------|
| **Granularity** | Single variable | Arbitrary code block |
| **Performance** | Fastest (CPU-level) | Slightly slower (OS scheduler) |
| **Use case** | Counters, flags, state | Complex data structures, multi-variable updates |
| **Blocking** | Non-blocking | Blocks goroutine |
| **Deadlock risk** | None | Possible if misused |

**Rule of thumb:**
- Single numeric counter/flag → **atomic**
- Multiple fields or complex logic → **mutex**
- Many readers, few writers → **RWMutex**

---

## Final Proverb (repeated for emphasis)

> **Don't communicate by sharing memory; share memory by communicating.**

Go gives you both low-level synchronisation (atomic, mutex) and high-level concurrency primitives (channels). Pick the right tool for the job.
