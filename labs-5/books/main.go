package books

import (
	"fmt"
)

func Run() {
	book := Book{
		Title:  "1984",
		Author: "George Orwell",
		Year:   1949,
	}

	fmt.Println(book)
}
