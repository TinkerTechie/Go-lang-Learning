// Package main defines this as an executable Go program.
package main

// We import two standard packages:
// fmt  -> for printing output to the console
// math -> for mathematical constants and functions
import (
	"fmt"
	"math"
)

// main is the entry point of the program.
func main() {

	// In Go, only names that start with a CAPITAL letter are exported.
	// Exported names can be accessed from other packages.

	// Pi is an EXPORTED name from the math package.
	// That is why we can access it as math.Pi
	fmt.Println(math.Pi)

	// If we tried to use math.pi (lowercase),
	// it would give a compilation error because:
	// - pi is NOT exported
	// - unexported names cannot be accessed outside their package
}
