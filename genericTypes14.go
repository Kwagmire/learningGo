package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	a, b := 1, 2
	//i, j := &a, &b

	var m List[int]
	n := List[int]{&m, a}
	m.val = b

	fmt.Println(n, n.next)
}
