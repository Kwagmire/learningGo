package main

import (
	"fmt"
)

func Tri(x int) {
	for i := 1; i <= x; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}
}

func main() {
	Tri(4)
}
