package example

import (
	"fmt"

	string_utils "labs_programming_go/labs-3/utils/string"
)

func exampleReverseString() {
	fmt.Println("\nReverse string example:")

	var input string
	fmt.Print("Введите строку для переворота: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	reversed := string_utils.Reverse(input)
	fmt.Printf("Перевернутая строка: %s\n", reversed)
}
