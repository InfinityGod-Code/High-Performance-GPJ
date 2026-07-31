package channel

import "fmt"

// ====================================================================
// CLOSING A CHANNEL — signals that no more values will be sent
// Ranging over a channel reads until it's closed.
// ====================================================================
func RangeCloseDemo() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch) // signals the receiver that we're done
	}()

	for val := range ch { // automatically stops when ch is closed
		fmt.Println("Ranged:", val)
	}

	// After close, the channel still yields zero values
	val, ok := <-ch
	fmt.Printf("After close: val=%d, ok=%t\n", val, ok) // ok=false means channel is closed
}
