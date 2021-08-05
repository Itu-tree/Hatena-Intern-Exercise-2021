package main

import (
	"fmt"
	"math"
)



func Sqrt(x float64) float64 {
	z := float64(1)
	d := 1.0
	for math.Abs(d) > 1e-10 {
		d = (z*z - x) / (2 * z)
		z -= d
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
