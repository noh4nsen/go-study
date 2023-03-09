package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // Ponteiro para i
	fmt.Println(*p) // Le i atraves do ponteiro
	*p = 21         // Altera i atraves do ponteiro
	fmt.Println(i)

	p = &j       // Ponteiro para j
	*p = *p / 37 // altera j atraves do ponteiro
	fmt.Println(j)
}
