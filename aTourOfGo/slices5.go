package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var s [][]uint8

	for i := 0; i < dy; i++ {
		var m []uint8
		s = append(s, m)
		for j := 0; j < dx; j++ {
			s[i] = append(s[i], uint8(i+j)/2)
		}
	}

	return s

}

func main() {

	//fmt.Println(Pic(2,3))
	pic.Show(Pic)
}
