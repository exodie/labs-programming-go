package books

import "fmt"

func (b Book) String() string {
	return fmt.Sprintf("Title: %s, Author: %s, Year: %d", b.Title, b.Author, b.Year)
}
