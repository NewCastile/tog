package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	// Just picking a color model for this example.
	return color.RGBAModel
}
func (i Image) Bounds() image.Rectangle {
	// Hardcoding this image is 16x16
	return image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: 16, Y: 16},
	}
}
func (img Image) At(x, y int) color.Color {
	// v := uint8(x * y)
	// Other suggested values were:
	v := uint8(x * y)
	// v := uint8((x+y)/2)
	return color.RGBA{v, v, 255, 255}
}
func main() {
	m := Image{}
	pic.ShowImage(m)
}
