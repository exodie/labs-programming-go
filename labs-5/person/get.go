package person

import "fmt"

func (p Person) Info() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.name, p.age)
}
