package main

import "fmt"

func main() {
	sum := 1

	// diretivas init (i := 0) e post (i++)
	// sao opcionais.
	// for tem mesmo funcionamento do while em Go
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}
