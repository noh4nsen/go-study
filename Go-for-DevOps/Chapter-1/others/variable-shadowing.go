package others

import "fmt"

var word = "hello"

func VariableShadowing() {
	var word = "world"

	fmt.Println("Inside main(): ", word)
	printOutter()
}

func printOutter() {
	fmt.Println("The package level 'word' var: ", word)
}
