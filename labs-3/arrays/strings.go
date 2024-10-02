package arrays

import (
	"fmt"
)

func maxLengthOfStrings() {
	fmt.Println("\nFind max length string of array:")

	strings := []string{"apple", "banana", "cherry", "blueberry", "very very very long banana=D"}

	var longestString string

	for _, str := range strings {
		if len(str) > len(longestString) {
			longestString = str
		}
	}

	fmt.Println("Самая длинная строка:", longestString)
}
