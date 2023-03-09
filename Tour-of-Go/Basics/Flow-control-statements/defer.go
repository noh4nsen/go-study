package main

import "fmt"

func main() {
	// defer atrasa o retorno ate que as outras funcoes retornem
	defer fmt.Println("world")
	fmt.Println("hello")
}
