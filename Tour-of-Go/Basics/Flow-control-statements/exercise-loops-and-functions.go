package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	for i := 0; i <= 100; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, " Inner: ", z)
		if z == math.Sqrt(x) {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println("func Sqrt: ", Sqrt(5))
	fmt.Println("math.Sqrt: ", math.Sqrt(5))
}
