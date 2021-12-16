// Go provides math library which you can read on https://golang.org/pkg/math;

package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2.4
	b := 1.6

	c := math.Pow(a, 2)
	fmt.Printf("%.2f^%d = %.2f \n", a, 2, c)

	d := math.Sin(a)
	fmt.Printf("Sin(%.2f) = %.2f \n", a, d)

	e = math.Cos(b)
	fmt.Printf("Cos(%.2f) = %.2f \n", b, e)

	f = math.Sqrt(a * b)
	fmt.Printf("Sqrt(%.2f*%.2f) = %.2f \n", a, b, f)
}
