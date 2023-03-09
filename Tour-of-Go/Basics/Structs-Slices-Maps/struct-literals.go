package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}  // Y eh zero implicitamente
	v3 = Vertex{}      // X e Y sao zero implicitamente
	p  = &Vertex{1, 2} // Tem como tipo *Vertex (ponteiro para o struct)
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
