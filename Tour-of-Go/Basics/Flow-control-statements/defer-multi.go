package main

import "fmt"

func main() {
	fmt.Println("counting")

	// chamadas defer sao empilhadas e chamadas
	// em lifo
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
