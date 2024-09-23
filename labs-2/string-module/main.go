package string_module

import "fmt"

func Run() {
	fmt.Println("\nString length:")

	var input = "hello world"

	// ! В задании не указано что строка должна вписываться из под терминала (scanf) !
	//fmt.Print("Введите строку: ")
	//_, err := fmt.Scan(&input)
	//if err != nil {
	//	fmt.Println("Ошибка ввода:", err)
	//	return
	//}

	fmt.Println("Длина строки:", length(input))
}
