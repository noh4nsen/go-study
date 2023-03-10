package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value: ", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value: ", m["Answer"])

	// v eh o elemento
	// ok eh true se a chave "Answer" existe no map
	// ok eh false se a chave "Answer" nao existe no map
	v, ok := m["Answer"]
	fmt.Println("The value: ", v, "Present?", ok)
}
