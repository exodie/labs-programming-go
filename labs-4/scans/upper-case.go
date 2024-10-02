package scans

import (
	"fmt"
	"strings"
)

func toUpperCase() {
	fmt.Println("\nUpper case to string:")

	var input string

	fmt.Print("Введите строку: ")
	fmt.Scanln(&input)

	upperInput := strings.ToUpper(input)

	fmt.Println(upperInput)
}
