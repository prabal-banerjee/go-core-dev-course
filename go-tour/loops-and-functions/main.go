package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	i := 1
	for delta := 1.0; math.Abs(delta) > 1e-10; {
		delta = (z*z - x) / (2 * z)
		z -= delta
		i += 1
	}
	fmt.Println("Number of iterations:", i)
	return z
}

func main() {
	x := 2.0
	fmt.Println("Newton method sqrt:", Sqrt(x))
	fmt.Println("Math lib sqrt:", math.Sqrt(x))
}
