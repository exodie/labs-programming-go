package number_module

import (
	"fmt"
)

func average(a int, b int) float64 {
	return float64(a+b) / 2.0
}

func averageOfTwoNumbers() {
	fmt.Println("\nAverage of two numbers:")

	var num1, num2 int

	fmt.Print("Введите первое целое число: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	fmt.Print("Введите второе целое число: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	avg := average(num1, num2)
	fmt.Printf("Среднее значение: %.2f\n", avg)
}
