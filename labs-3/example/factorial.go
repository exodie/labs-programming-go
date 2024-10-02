package example

import (
	"fmt"
	"labs_programming_go/labs-3/utils/math"
)

func exampleFactorial() {
	fmt.Println("Factorial example:")

	var number int
	fmt.Print("Введите число для вычисления факториала: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	factorial := math.Factorial(number)
	fmt.Printf("Факториал %d равен %d\n", number, factorial)
}
