package main

import (
	"fmt"
	"math"
)

// define a struct
// Circle struct has three fields: x,y and r
// Basically, we can create struct method in the same way when you create a function
type Circle struct {
	x, y int
	r    float64
}

func (c Circle) display() {
	fmt.Printf("x=%d, y=%d, r=%.2f \n", c.x, c.y, c.r)
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (c Circle) circumference() float64 {
	return 2 * math.Pi * c.r
}

func (c Circle) moveTo(newX, newY int) {
	c.x = newX
	c.y = newY
}

func main() {
	// methods
	shape := Circle{10, 5, 2.8}
	shape.display()
	fmt.Printf("area = %2.f \n", shape.area())
	fmt.Printf("circumference=%2.f\n", shape.circumference())

	shape.moveTo(5, 5)
	shape.display()
}
