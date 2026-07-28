package main

import (
	"log"
	"sync"
	"fmt"
)
 
/*
Running without the Go..
*/
func main(){
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
}