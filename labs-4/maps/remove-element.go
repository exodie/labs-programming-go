package maps

import (
	"fmt"
)

func removeEntryByName(data map[string]string, name string) {
	if _, exists := data[name]; exists {
		delete(data, name)
		fmt.Printf("Запись с именем '%s' была удалена.\n", name)
	} else {
		fmt.Printf("Запись с именем '%s' не найдена.\n", name)
	}
}

func RunExampleRemoveElementOfMaps() {
	fmt.Println("\nRemove element of maps:")

	data := map[string]string{
		"Alice":   "alice@example.com",
		"Bob":     "bob@example.com",
		"Charlie": "charlie@example.com",
	}

	var nameToRemove string

	for nameOfData := range data {
		fmt.Println(nameOfData, " ")
	}

	fmt.Print("Введите имя для удаления: ")
	fmt.Scanln(&nameToRemove)

	removeEntryByName(data, nameToRemove)

	fmt.Println("Оставшиеся записи:")
	for name, email := range data {
		fmt.Printf("%s: %s\n", name, email)
	}
}
