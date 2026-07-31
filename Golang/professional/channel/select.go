package channel

import (
	"fmt"
	"time"
)

// ====================================================================
// SELECT — waits on multiple channel operations simultaneously
// The first ready case is chosen at random (fair scheduling).
// ====================================================================
func SelectDemo() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	// Wait for whichever arrives first (ch1 will win since it sleeps less)
	select {
	case msg := <-ch1:
		fmt.Println("Select won:", msg)
	case msg := <-ch2:
		fmt.Println("Select won:", msg)
	}
}

// ====================================================================
// SELECT WITH TIMEOUT — prevents waiting forever
// ====================================================================
func SelectWithTimeout() {
	ch := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "result"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Got result:", msg)
	case <-time.After(100 * time.Millisecond): // fires before the goroutine finishes
		fmt.Println("Timeout: operation took too long")
	}
}

// ====================================================================
// SELECT WITH DEFAULT — non-blocking channel operation
// ====================================================================
func SelectWithDefault() {
	ch := make(chan int, 1)
	ch <- 10

	// Non-blocking read
	select {
	case val := <-ch:
		fmt.Println("Read:", val)
	default:
		fmt.Println("No value available (won't print since we put one)")
	}

	// Non-blocking read from empty channel
	select {
	case val := <-ch:
		fmt.Println("Read:", val)
	default:
		fmt.Println("No value available (channel is empty)")
	}
}

// ====================================================================
// Runs all select examples
// ====================================================================
func SelectAll() {
	fmt.Println("--- Select Demo ---")
	SelectDemo()

	fmt.Println("\n--- Select With Timeout ---")
	SelectWithTimeout()

	fmt.Println("\n--- Select With Default ---")
	SelectWithDefault()
}
