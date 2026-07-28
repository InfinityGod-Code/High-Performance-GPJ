// A special package that tells the Go compiler to build an executable program.
package main

import (
	"fmt"
)

/* Within the main package there must be main() function, which acts as the entry point of your program */

func main(){
	fmt.Println(name)
	fmt.Println(DEBUG, LogLevel, timeStamp)

	// accessing mutiple variables from the function 
	age, name := shortHandDeclarations()
	fmt.Println(age, name)
	pointers()
}