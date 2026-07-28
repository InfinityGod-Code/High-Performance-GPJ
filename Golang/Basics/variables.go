package main

import (
	"time"
)

//multiple variable declaration in Go
var (
	DEBUG bool = false
	LogLevel string = "info"
	timeStamp time.Time = time.Now()
)


// Short variable declaration : Instead if writing var name type we can use := operator for declaring 
// variables in go. 

/* 
So if we have : 
var age int = 25
var name string = "Alice"
In Short variable declaration can be written as : 
*/

/*
Scope : 
:= can only be used inside functions.
Using this at package level will be error 
At package level we must be using the var
See How this declaring multiple variables from a function accessed from the main.
*/
func shortHandDeclarations() (int,string) {
	age := 25
    name := "Alice"
	return age,name
}



