## Quick Guide 

#### Running Instructions 
- There project folder structure in Go goes like : 
  Package 
  Module

- Package : A package is Go's way of organizing related code files within the same directory. Every .go file must start with a package declaration.
  
  - Single Directory: All files in a single folder must belong to the same package.
  - Visibility Rules: Go uses capitalization to control access (no public or private keywords).Exported (Public):    Starts with a Capital letter (e.g., fmt.Println, math.Pi). Accessible from other packages.Unexported (Private): Starts with a lowercase letter (e.g., calculateSum). Only accessible within its own package.

  - The main Package: A special package that tells the Go compiler to build an executable program rather than a shared library. It must contain a main() function, which acts as the entry point of your program.

- Module : A module is a collection of related Go packages that are versioned, released, and distributed together. It tracks dependencies.

  - The go.mod File: Found at the root directory of your project. It defines the module's import path and lists the exact versions of third-party libraries your project needs.

  - In order to compile and run the entire module and execute the main package we can use command : 
    ** go run . ** 