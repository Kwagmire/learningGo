package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	prev := 0.15
	for math.Abs(prev-z) > 0.000001 {
		prev = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	Sqrt(144)
	fmt.Printf("\n%v", math.Sqrt(144))
}
