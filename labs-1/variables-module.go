package labs_1

import "fmt"

func variablesModule() {
	fmt.Println("\nVariables module:")

	var number int = 42                  // number := 42
	var floatNumber float64 = 3.14       // floatNumber := 3.14
	var desc string = "This is a string" // desc := "This is a string"
	var status bool = true               // status := true

	fmt.Println("Целое число (int):", number)
	fmt.Println("Число с плавающей точкой (float64):", floatNumber)
	fmt.Println("Строка (string):", desc)
	fmt.Println("Логическое значение (bool):", status)
}
