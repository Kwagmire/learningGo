package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
}

func (i Image) ColorModel() color.Model { return color.RGBAModel }

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{500, 500}
	pic.ShowImage(m)
}
