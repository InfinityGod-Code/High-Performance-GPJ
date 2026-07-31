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
6. [Channels](#6-channels)

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

**Reference:** `mutex.go`

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

---

## 6. Channels

> "Don't communicate by sharing memory; share memory by communicating."

Channels are Go's **native way to communicate between goroutines**. A channel is a typed conduit — you send values of a specific type into it and receive values of that type from it.

**Reference:** `channel/` directory

### 6.1 Basic Send & Receive

**Reference:** `channel/basics.go`

```go
ch := make(chan int)   // unbuffered channel of int
go func() { ch <- 1 }()  // send (blocks until received)
val := <-ch              // receive (blocks until sent)
```

- By default, sends **block** until a receiver is ready and vice versa — this is how goroutines synchronise without explicit locks or condition variables.

### 6.2 Unbuffered vs Buffered Channels

**Reference:** `channel/buffered.go`

| Type | Creation | Behaviour |
|------|----------|-----------|
| **Unbuffered** | `make(chan T)` | Sender blocks until receiver takes the value |
| **Buffered** | `make(chan T, N)` | Sender blocks only when buffer (size N) is full |

```go
ch := make(chan int, 3)  // buffer of 3
ch <- 1; ch <- 2; ch <- 3  // all succeed immediately
// ch <- 4  // would block (or deadlock)
```

- `len(ch)` — number of elements currently buffered
- `cap(ch)` — capacity of the buffer

### 6.3 Directional Channels

**Reference:** `channel/directional.go`

Channel types can be constrained to send-only or receive-only as function parameters. This is enforced at **compile time** and makes APIs safer.

```go
func SendOnly(ch chan<- int, v int) { ch <- v }    // can only send
func ReceiveOnly(ch <-chan int) { v := <-ch }      // can only receive

ch := make(chan int)   // bidirectional at creation
go SendOnly(ch, 42)
ReceiveOnly(ch)        // callers pass the same channel; the callee's view is restricted
```

### 6.4 Closing & Ranging Over Channels

**Reference:** `channel/range_close.go`

- A sender calls `close(ch)` to signal that no more values will be sent.
- Receivers can use `for v := range ch` to read until the channel is closed.
- The two-value receive `v, ok := <-ch` yields `ok=false` when the channel is closed.

```go
ch := make(chan int)
go func() {
    for i := 0; i < 5; i++ { ch <- i }
    close(ch)
}()
for v := range ch { fmt.Println(v) }  // stops when closed
```

> **Important:** Only the **sender** should close a channel. Closing a channel twice or sending on a closed channel causes a panic.

### 6.5 Select Statement

**Reference:** `channel/select.go`

`select` lets a goroutine **wait on multiple channel operations** simultaneously. The first case that becomes ready is chosen at random (fair scheduling).

```go
select {
case msg := <-ch1:
    fmt.Println("from ch1:", msg)
case msg := <-ch2:
    fmt.Println("from ch2:", msg)
case <-time.After(100 * time.Millisecond):
    fmt.Println("timeout")
default:
    fmt.Println("non-blocking: no channel ready")
}
```

- `time.After` creates a channel that fires after a duration — useful for timeouts.
- `default` makes a select **non-blocking**; if no case is ready, default runs immediately.
- An empty `select {}` blocks forever.

### 6.6 Worker Pool (Fan-Out / Fan-In)

**Reference:** `channel/worker_pool.go`

The **worker pool** pattern dispatches jobs to multiple goroutines (fan-out) and collects results (fan-in).

```
         ┌─────────┐
 Jobs ──▶│ Worker 1│──┐
         ├─────────┤  │  ┌──────────┐
         │ Worker 2│──┼─▶│ Results  │
         ├─────────┤  │  └──────────┘
         │ Worker 3│──┘
         └─────────┘
```

```go
const numWorkers = 3
jobs := make(chan int, 10)
results := make(chan int, 10)
var wg sync.WaitGroup

for i := 1; i <= numWorkers; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        for job := range jobs {
            results <- job * 2  // do work
        }
    }(i)
}

for j := 1; j <= 10; j++ { jobs <- j }
close(jobs)       // signal workers: no more jobs
wg.Wait()         // wait for all workers
close(results)    // now safe to range over results
for r := range results { fmt.Println(r) }
```

Key points:
- `close(jobs)` tells workers to stop ranging.
- `wg.Wait()` ensures all workers finish before closing the results channel.
- Results must be closed **after** all workers are done, otherwise ranging over results deadlocks.

---

## Final Proverb (repeated for emphasis)

> **Don't communicate by sharing memory; share memory by communicating.**

Go gives you both low-level synchronisation (atomic, mutex) and high-level concurrency primitives (channels). Pick the right tool for the job.
