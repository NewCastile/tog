package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	outer := make([][]uint8, dy)
	
	for i := range outer {
	 	inner := make([]uint8, dx)
		
		for j := range inner {
			inner[j] = uint8((j^i) * (i^j))
		}
		
		outer[i] = inner
	}
	
	return outer
}

func main() {
	pic.Show(Pic)
}