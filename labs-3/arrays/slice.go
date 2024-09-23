package arrays

import (
	"fmt"
)

func sliceArraysNumbers() {
	fmt.Println("\nSlice array numbers:")

	numbers := [5]int{1, 2, 3, 4, 5}

	slice := numbers[1:4]

	fmt.Println("Исходный срез:", slice)

	slice = append(slice, 6)
	fmt.Println("Срез после добавления 6:", slice)

	indexToRemove := 2
	if indexToRemove < len(slice) {
		slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
		fmt.Println("Срез после удаления элемента по индексу 2:", slice)
	} else {
		fmt.Println("Индекс для удаления выходит за пределы среза.")
	}
}
