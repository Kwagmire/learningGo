package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	sum := 0
	previous := -1
	before := 0
	return func() int {
		if previous == -1 {
			sum = 0
			previous = 0
		} else if previous == 0 {
			sum = 1
			previous = 1
		} else {
			sum = previous + before
			before = previous
			previous = sum
		}
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
