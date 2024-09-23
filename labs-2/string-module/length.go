package string_module

func length(str string) int {
	return len([]rune(str))
}
