package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	for i := 0; i <= 3; i++ {
		defer fmt.Print(i)
	}
}
