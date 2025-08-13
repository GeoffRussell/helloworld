package main

import (
	"fmt"
	"math"

	"github.com/GeoffRussell/helloWorld/morestrings"
	"github.com/google/go-cmp/cmp"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for math.Abs(z*z-x) > 0.000000002 {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	s := Sqrt(2)
	fmt.Println("Sqrt of 2:  ", s, " check s*s: ", s*s)
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println("Hello, world \u00E9 é")
	fmt.Println(cmp.Diff("Hello, world\nXXX", "Hello, Go\nYYY"))
}
