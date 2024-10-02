package maps

import (
	"fmt"
)

func createMapsWithPeople() {
	fmt.Println("\nCreating maps with people:")

	people := make(map[string]int)

	people["Alice"] = 30
	people["Bob"] = 25
	people["Charlie"] = 35

	addPerson := func(name string, age int) {
		people[name] = age
	}

	addPerson("David", 28)

	fmt.Println("Список людей и их возрастов:")
	for name, age := range people {
		fmt.Printf("Имя: %s, Возраст: %d\n", name, age)
	}
}
