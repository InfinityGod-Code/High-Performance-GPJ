package channel

import "fmt"

// ====================================================================
// SEND-ONLY — this function can only SEND to the channel
// chan<-  means data flows OUT of this function (write)
// ====================================================================
func SendOnly(sendCh chan<- int, value int) {
	sendCh <- value
	// readCh := <- sendCh // COMPILE ERROR: cannot receive from send-only channel
}

// ====================================================================
// RECEIVE-ONLY — this function can only RECEIVE from the channel
// <-chan  means data flows INTO this function (read)
// ====================================================================
func ReceiveOnly(receiveCh <-chan int) {
	val := <-receiveCh
	// receiveCh <- 99 // COMPILE ERROR: cannot send to receive-only channel
	fmt.Println("Directional: received", val)
}

// ====================================================================
// Demos wiring a send-only and receive-only channel together
// ====================================================================
func DirectionalDemo() {
	ch := make(chan int) // bidirectional channel

	go SendOnly(ch, 100) // SendOnly restricts its view to send-only
	ReceiveOnly(ch)      // ReceiveOnly restricts its view to receive-only
}
