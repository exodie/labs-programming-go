package maths

import (
	"fmt"
	"math"
)

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func printAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("Площади: %.2f\n", shape.area())
	}
}
