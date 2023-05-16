package pointers

func ChangeValue(word *string) {
	*word += "world"
}
