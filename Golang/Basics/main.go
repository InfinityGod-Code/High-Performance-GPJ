// A special package that tells the Go compiler to build an executable program.
package main

import (
	"fmt"
)

// Declaring variables in GO involve 
// Single varible creation in Go
var name string = "InfinityGod-Code"
 

/* Within the main package there must be main() function, which acts as the entry point of your program */

func main(){
	fmt.Println(name)
	fmt.Println(DEBUG, LogLevel, timeStamp)
}