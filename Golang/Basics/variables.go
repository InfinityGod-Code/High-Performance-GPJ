package main

import (
	"time"
)

// Declaring variables in GO involve
// Single varible creation in Go
var name string = "InfinityGod-Code"

// multiple variable declaration in Go
var (
	DEBUG     bool      = false
	LogLevel  string    = "info"
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
func shortHandDeclarations() (int, string) {
	age := 25
	name := "Alice"
	return age, name
}

/*
Constants and Enums in Go
*/

// Constants: immutable values declared with const keyword
const AppVersion = "1.0"       // untyped constant (inferred type)
const MaxConnections int = 100 // typed constant

// Grouped constant declaration
const (
	StatusOK       = 200
	StatusNotFound = 404
)

// Enum pattern using iota (auto-incrementing constant generator)
const (
	_      = iota
	Low    // 1
	Medium // 2
	High   // 3
)
