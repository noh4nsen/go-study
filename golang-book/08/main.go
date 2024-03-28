package main

import "fmt"

func divide(nominator int, divider int) float32 {
	if divider == 0 {
		panic("can't divide by 0")
	}
	return float32(nominator) / float32(divider)
}

func main() {
	fmt.Println(divide(3, 2))
}
