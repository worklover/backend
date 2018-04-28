package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; math.Abs((z*z-x)/(2*z)) > 0.000000000000001; i++ {
		z = z - (z*z-x)/(2*z)
		fmt.Println(i)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
