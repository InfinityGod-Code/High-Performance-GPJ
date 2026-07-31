package main

import (
	"fmt"
	"professional/channel"
	"sync"
	"time"
)

/*
Running without the Go..
*/
func main() {
	/*
		WaitGroups Demo =============================================
	*/
	// passing reference of this variable
	res := 0

	// passing reference of the WaitGroup to execute functions afterwards as well
	// this is the struct type
	wq := &sync.WaitGroup{}

	// I am about to start 1 goroutine, and you must wait for it to finish.
	wq.Add(1)
	fmt.Println("Inside the main function..")
	// calling goroutine with go keyword
	// Execute this function concurrently in a new goroutine.
	go sumWithWaitGroups(1, 5, wq, &res)
	wq.Wait()
	fmt.Println("Result:", res)

	fmt.Println("===================== Race Conditions Section ===================")

	count := 0

	fmt.Println(count) // this gives me 3 which is expected

	// Now trying the same with goroutine
	go raceConditions(&count)
	go raceConditions(&count)
	go raceConditions(&count)
	fmt.Println(count) // this print zero because the value of count might be 0 for all functions running in independent Goroutines:

	fmt.Println("===================== Channel Section ===================")

	// make create channel of type int
	ch := make(chan int)
	go channel.Greet(ch)

	// read value from the channel
	// while reading from channel, tail will be in front of channel
	temp := <-ch
	fmt.Println(temp)

	// Channel: Buffered vs Unbuffered
	channel.BufferedChannelDemo()

	// Channel: Directional (send-only / receive-only)
	channel.DirectionalDemo()

	// Channel: Range and Close
	channel.RangeCloseDemo()

	// Channel: Select
	channel.SelectAll()

	// Channel: Worker Pool (fan-out / fan-in)
	channel.WorkerPoolDemo()

	time.Sleep(100 * time.Millisecond) // allow goroutines to finish printing
}
