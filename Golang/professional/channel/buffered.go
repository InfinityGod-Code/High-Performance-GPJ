package channel

import "fmt"

// ====================================================================
// UNBUFFERED CHANNEL — blocks sender until receiver is ready
// ====================================================================
func UnbufferedDemo() {
	ch := make(chan int)

	go func() {
		ch <- 42 // blocks until main goroutine receives
	}()

	val := <-ch // receiver pulls the value
	fmt.Println("Unbuffered: received", val)
}

// ====================================================================
// BUFFERED CHANNEL — sender only blocks when buffer is full
// ====================================================================
func BufferedDemo() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Buffered: len =", len(ch), "cap =", cap(ch))

	// Sending a 4th value would block (or deadlock) because the buffer (size 3) is full.
	// Uncommenting the next line would cause: fatal error: all goroutines are asleep - deadlock!
	// ch <- 4

	// Reading all values drains the buffer
	fmt.Println("Buffered received:", <-ch)
	fmt.Println("Buffered received:", <-ch)
	fmt.Println("Buffered received:", <-ch)
}

// ====================================================================
// UNBUFFERED vs BUFFERED — visual comparison
// ====================================================================
func BufferedChannelDemo() {
	fmt.Println("--- Unbuffered Channel ---")
	UnbufferedDemo()

	fmt.Println("\n--- Buffered Channel ---")
	BufferedDemo()
}
