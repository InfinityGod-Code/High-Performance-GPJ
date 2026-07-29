package main

/* This section will help us understand how concurrent goroutines accessing shared data can lead to race conditions. 
We’ll explore how these issues occur and learn how Go’s synchronisation mechanisms can help us
safely coordinate goroutines and prevent data races.
*/

func raceConditions(templeBellCount *int) {
	// let's take the variable and increment unity each time it is called
	// deference first to perform any action 
	c := *templeBellCount
	*templeBellCount = c + 1
}

// if we call this function normally we can get the expected result 
// But trying with the goroutines that runs on completely different goroutine and does not aware of each other.

