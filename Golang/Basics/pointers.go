package main

import (
	"fmt"
)

// A pointer is a variable that stores the memory address of another value, rather than storing the value itself.

/*
You only need to know two operators to master Go pointers
:& (Ampersand / "Address of"): Finds the memory location of a variable.
* (Asterisk / "Value at / Dereference"): Looks inside that memory location to read or modify the value.
*/

func pointers(){
    // 1. Create a regular variable
    age := 25 

    // 2. Create a pointer using & (stores the memory address of age)
    var agePointer *int = &age 

    fmt.Println(age)        // Prints: 25
    fmt.Println(agePointer) // Prints a memory address, like: 0xc0000140e8

    // 3. Change the value using * (dereferencing)
    *agePointer = 30 

    fmt.Println(age)  
}  