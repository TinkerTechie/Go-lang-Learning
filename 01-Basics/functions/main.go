// Package main makes this file an executable Go program.
package main

// fmt package is imported to allow printing output to the console.
import "fmt"

// func is the keyword used to define a function in Go.
//
// add is the function name.
// (x int, y int) are parameters of the function.
// int at the end is the return type of the function.
//
// In Go, the type always comes AFTER the variable name.
func add(x int, y int) int {
	// The return statement sends the result back to the caller.
	return x + y
}

// main is the entry point of the program.
func main() {

	// Here we call the add function with two arguments: 42 and 13.
	// The returned result is printed using fmt.Println.
	fmt.Println(add(42, 13))
}
