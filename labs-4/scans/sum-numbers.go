package scans

import (
	"fmt"
)

func sumNumbers() {
	var sum float64
	var number float64

	fmt.Println("Введите числа (введите 'exit' для завершения):")

	for {
		var input string
		fmt.Scan(&input)

		if input == "exit" {
			break
		}

		_, err := fmt.Sscan(input, &number)
		if err != nil {
			fmt.Println("Пожалуйста, введите корректное число или 'exit' для завершения.")
			continue
		}

		sum += number
	}

	fmt.Printf("Сумма введенных чисел: %.2f\n", sum)
}
