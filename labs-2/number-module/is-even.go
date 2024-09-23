package number_module

import "fmt"

func isEven() {
	fmt.Println("\nNumber is even:")

	var number int

	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	if number%2 == 0 {
		fmt.Println("Число четное.")
	} else {
		fmt.Println("Число нечетное.")
	}
}
