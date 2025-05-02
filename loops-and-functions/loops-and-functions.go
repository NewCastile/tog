package main

import (
	"fmt"
	"math"
)

func computeSqrt(z, x float64) float64 {
	z -= (z*z - x) / (2 * z)
	return z
}

func Sqrt(x float64) float64 {
	i := 1
	val := 0.5
	inc := 0.05
	for i <= 40 {
		res := computeSqrt(val, x)
		fmt.Println("i:", i, "val", val, "res:", res, "res**", res*res)
		val = val + inc
		i++
	}
	return val
}

func SqrtTwo(x float64) float64 {
	fmt.Println("SqrtTwo(x)", x)

	val := 0.5
	const inc = 0.05 
	
	for i := inc; i <= 2; i += inc {
		res := computeSqrt(val, x)
		fmt.Println("i:", i, "val", val, "res:", res, "res**", res*res)
		val = val + inc
	}
	
	return val
}

func SqrtThree(x float64) float64 {
	fmt.Println("SqrtThree(x)", x)

	const inc = 0.05
	val := 0.5
	res := 10.0

	for computeSqrt(val, x) <= res {
		res = computeSqrt(val, x)
		fmt.Println("val", val, "res:", res, "res**", res*res)
		val += inc
	}

	return res
}

// Does not work
func SqrtFour(x float64) float64 {
	fmt.Println("SqrtFour(x)", x)

	const inc = 0.05
	val := 0.5
	res := 10.0

	for c := computeSqrt(val, x); c <= res; val += inc {
		res = computeSqrt(val, x)
		fmt.Println("val", val, "res:", res, "res**", res*res)
	}

	return res
}

// Does work
func SqrtFive(x float64) float64 {
	fmt.Println("SqrtFive(x)", x)

	const inc = 0.05
	val := 0.5
	res := 10.0

	for ;computeSqrt(val, x) <= res; val += inc {
		res = computeSqrt(val, x)
		fmt.Printf("val %f | res %f | res**%f\n", val, res, res*res)
	}

	return res
}


func main() {
	fmt.Println(SqrtFive(2))
	fmt.Println(math.Sqrt(2))
}
