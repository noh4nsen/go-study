package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// Primeiro elemento (i) eh o index
	// Segundo elemento (v) eh uma copia do elemento no index
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
