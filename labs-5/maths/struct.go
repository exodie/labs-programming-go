package maths

type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}
