package main

import "fmt"

func main() {
	var i, j int = 1, 2

	//Dentro de uma funcao podemos usar := ao inves de var
	//Fora de uma funcao necessariamente deve ser usado var
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
