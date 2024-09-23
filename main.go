package main

import (
	"fmt"
	labs_1 "labs_programming_go/labs-1"
	labs_2 "labs_programming_go/labs-2"
	labs_3 "labs_programming_go/labs-3"
)

func main() {
	fmt.Println("main pkg started...")

	var choice int
	fmt.Println("Выберите лабораторную работу для просмотра:")
	fmt.Println("1 - Лабораторная №1")
	fmt.Println("2 - Лабораторная №2")
	fmt.Println("3 - Лабораторная №3")
	fmt.Print("Введите номер лабораторной работы: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	switch choice {
	case 1:
		fmt.Println("\nЛабораторная №1:")
		labs_1.Run()
	case 2:
		fmt.Println("\nЛабораторная №2:")
		labs_2.Run()
	case 3:
		fmt.Println("\nЛабораторная №3:")
		labs_3.Run()
	default:
		fmt.Println("Неверный выбор. Пожалуйста, выберите 1 или 2.")
	}
}
