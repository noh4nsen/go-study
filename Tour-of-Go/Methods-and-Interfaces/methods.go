package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Metodo de Vertex, tem o receiver (v Vertex)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Funcao normal
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	vert := Vertex{3, 4}
	fmt.Println(vert.Abs())
	fmt.Println(Abs(vert))
}
