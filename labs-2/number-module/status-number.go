package number_module

import "fmt"

func statusAsNumber() {
	fmt.Println("\nStatus as number:")

	var number int

	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	var result string

	if number > 0 {
		result = "Positive"
	} else if number < 0 {
		result = "Negative"
	} else {
		result = "Zero"
	}

	fmt.Println("Результат:", result)
}
