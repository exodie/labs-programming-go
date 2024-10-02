package scans

import "fmt"

func revArray() {
	fmt.Println("\nReverse of array numbers:")

	var n int

	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)

	numbers := make([]int, n)

	fmt.Println("Введите числа:")
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	fmt.Println("Числа в обратном порядке:")
	for i := n - 1; i >= 0; i-- {
		fmt.Print(numbers[i], " ")
	}
	fmt.Println()
}
