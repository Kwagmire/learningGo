package main

import (
	"strings"

	"golang.org/x/tour/wc"
	//"fmt"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	for _, v := range strings.Fields(s) {
		_, ok := m[v]
		if ok == true {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}

	return m
}

func main() {
	//fmt.Println(WordCount("i am that i am"))
	wc.Test(WordCount)
}
