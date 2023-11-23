package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// Declare and initialize two integer variables, x and y, with values 3 and 4.

	var f float64 = math.Sqrt(float64(x*x + y*y))
	// Calculate the square root of the sum of the squares of x and y.
	// The result is assigned to a float64 variable f.

	var z uint = uint(f)
	// Convert the float64 value f to an unsigned integer (uint) and assign it to z.

	fmt.Println(x, y, z)
	// Print the values of x, y, and z.
}
