package main

/*
 In Go main function works as GoRoutine itself
 Let's see this with the example :
 let say we launch another goroutine inside this main goroutine function
 simple goroutine function will not even run and main gorotine can finish and exits
 So if go main goroutine exits it will also not run any existing goroutine that was
 running.

*/

import (
	"fmt"
	"sync"
)

func sum(from, to int) int {
	total := 0

	for i := from; i <= to; i++ {
		total += i
	}
	message := fmt.Sprintf("Inside the sum : %d", total)
	fmt.Println(message)
	return total
}

// WaitGroups are passed as reference therefore we can control from outside as well
// by calling wq.Wait() function

// Note : Generally for performance we use pass by reference instead pass by copy
// instead of running value we pass variable as reference as dereference it later use.
func sumWithWaitGroups(from, to int, wq *sync.WaitGroup, res *int) {
	*res = 0
	for i := from; i <= to; i++ {
		*res += i
	}

	// this will tell the WaitGroup that
	wq.Done()
}

/*
The WaitGroup establishes the necessary synchronization.
If multiple goroutines were writing to res concurrently, you'd need to think carefully about data races
and probably use a mutex, atomic operation, or—often more idiomatically for results—a channel.

Fortunately we have something called the Atomic which helps us safely overrides the single value accross different
goroutines
*/
