package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Metodos com ponteiros alteram o valor da instancia do objeto
// sem o ponteiro, a operacao eh feita sobre a copia
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Resultado equivalente ao metodo que usa ponteiro
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}

	v.Scale(10)
	fmt.Println(v.Abs())

	v.Scale(20)
	fmt.Println(v.Abs())

	v.Scale2(10)
	fmt.Println(v.Abs())

	Scale(&v, 5)
	fmt.Println(v.Abs())

}
