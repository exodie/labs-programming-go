package labs_1

import "fmt"

func averageThreeNumbers(a float64, b float64, c float64) float64 {
	var result = (a + b + c) / 3

	return result
}

func calculateFloatingNumber(a float64, b float64) (float64, float64) {
	sum := a + b
	difference := a - b
	return sum, difference
}

func arithmeticOperationsModule() {
	fmt.Println("\nArithmetic operations module:")

	var num1, num2 float64

	fmt.Print("Введите первое число с плавающей запятой (или нажмите Enter для значения по умолчанию 4): ")
	var input1 string
	fmt.Scanln(&input1)
	if input1 != "" {
		fmt.Sscan(input1, &num1)
	} else {
		num1 = 4
	}

	fmt.Print("Введите второе число с плавающей запятой (или нажмите Enter для значения по умолчанию 2): ")
	var input2 string
	fmt.Scanln(&input2)
	if input2 != "" {
		fmt.Sscan(input2, &num2)
	} else {
		num2 = 2
	}

	sum, difference := calculateFloatingNumber(num1, num2)

	product := num1 * num2
	quotient := num1 / num2

	fmt.Println("Сумма:", sum)
	fmt.Println("Разность:", difference)
	fmt.Println("Произведение:", product)
	fmt.Println("Частное:", quotient)

	fmt.Println("Вычисление среднего значения трех чисел:")

	num1 = 0
	num2 = 0
	var num3 float64

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&num1)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&num2)
	fmt.Print("Введите третье число: ")
	fmt.Scanln(&num3)

	average := averageThreeNumbers(num1, num2, num3)
	fmt.Println("Среднее значение:", average)
}
