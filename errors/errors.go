package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < float64(0) {
		e := ErrNegativeSqrt(x)
		return 0, e
	}
	const inc = 0.05
	val := 0.5
	res := 10.0

	for ; computeSqrt(val, x) <= res; val += inc {
		res = computeSqrt(val, x)
		fmt.Printf("val %f | res %f | res**%f\n", val, res, res*res)
	}

	return res, nil
}

func computeSqrt(z, x float64) float64 {
	z -= (z*z - x) / (2 * z)
	return z
}

func main() {
	sqrt, sqrtErr := Sqrt(2.5)
	fmt.Println(sqrt, sqrtErr)
	nsqrt, nsqrtErr := Sqrt(-2.5)
	fmt.Println(nsqrt, nsqrtErr)
}