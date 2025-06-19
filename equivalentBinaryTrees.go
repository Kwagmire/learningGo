package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {return}
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
 		Walk(t1, ch1)
		defer close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		defer close(ch2)
	}()
	for {
		x, ok1 := <-ch1
		y, ok2 := <-ch2
		
		//fmt.Println(x,y)
		//fmt.Println(ok1,ok2)
		
		if x != y || ok1 != ok2 {return false}
		if ok1 == false && ok1 == ok2 {return true}
	}
}

func main() {
	//Same(tree.New(1), tree.New(1))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
