package arrays

import (
	"fmt"
)

func arrayFiveNumbers() {
	fmt.Println("\nArray of five numbers:")

	var numbers [5]int

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Введите число %d: ", i+1)
		_, err := fmt.Scan(&numbers[i])
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			return
		}
	}

	fmt.Println("Введенные числа:")
	for _, number := range numbers {
		fmt.Println(number)
	}
}
