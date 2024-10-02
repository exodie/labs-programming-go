package person

import "fmt"

func Run() {
	person := Person{name: "Alice", age: 30}

	fmt.Println(person.Info())

	person.birthday()

	fmt.Printf("After birthday: %s is %d years old.\n", person.name, person.age)
}
