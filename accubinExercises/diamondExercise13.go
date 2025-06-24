package main

import (
	"fmt"
)

func Tri(x int) {
	num := 1
	for i := 1; i <= x; i++ {
		spaces := x - i
		for m := 0; m < spaces; m++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("%v ", num)
		}
		for m := 0; m < spaces; m++ {
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
		num++
	}

	num -= 2
	for i := x - 1; i > 0; i-- {
		spaces := x - i
		for m := 0; m < spaces; m++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("%v ", num)
		}
		for m := 0; m < spaces; m++ {
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
		num--
	}
}

func main() {
	Tri(20)
}
