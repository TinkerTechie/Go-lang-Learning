// -----------------------------
// PACKAGE DECLARATION
// -----------------------------

// package main defines the entry point of a Go program.
// Only files in package "main" can produce an executable.
package main

// -----------------------------
// IMPORT STATEMENTS
// -----------------------------

// import is used to include other packages so we can reuse their code.
// fmt       -> provides formatted input/output functions (printing to console)
// math/rand -> provides functions to generate pseudo-random numbers
import (
	"fmt"
	"math/rand"
)

// -----------------------------
// MAIN FUNCTION
// -----------------------------

// func is the keyword used to define a function in Go.
// main is a special function name.
// Execution of the program always starts from func main().
func main() {

	// fmt.Println prints text to the console.
	// It automatically adds a space between arguments
	// and moves the cursor to the next line.
	//
	// rand.Intn(10) generates a random integer
	// between 0 (inclusive) and 10 (exclusive).
	fmt.Println("My favourite number is", rand.Intn(10))
}
