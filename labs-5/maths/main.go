package maths

import (
	"fmt"
)

func Run() {
	circle := Circle{radius: 5.0}
	rectangle := Rectangle{width: 10.5, height: 8.5}

	area := circle.area()

	fmt.Printf("The area of the maths with radius %.2f is %.2f.\n", circle.radius, area)

	fmt.Println("\nВычисление всех площадей:")
	shapes := []Shape{circle, rectangle}

	printAreas(shapes)
}
