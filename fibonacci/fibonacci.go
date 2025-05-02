package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.

type FibonacciValue struct {
	X int
	Y int
}

// Take index(i) and create a pair of values with p := { x: i, y: i+1 }

func fibonacci() func() int { 
	fibonacciMap := FibonacciValue{ X: 0, Y: 1 }
	return func() int {
		fibonacciMap.X = fibonacciMap.Y
		fibonacciMap.Y = fibonacciMap.X + fibonacciMap.Y
		return fibonacciMap.Y
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}