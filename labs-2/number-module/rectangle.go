package number_module

import (
	"fmt"
)

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func rectangleSize() {
	fmt.Println("\nRectangle size:")

	var width, height float64

	fmt.Print("Введите ширину прямоугольника: ")
	_, err := fmt.Scan(&width)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	fmt.Print("Введите высоту прямоугольника: ")
	_, err = fmt.Scan(&height)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	rect := Rectangle{Width: width, Height: height}

	area := rect.Area()
	fmt.Printf("Площадь прямоугольника: %.2f\n", area)
}
