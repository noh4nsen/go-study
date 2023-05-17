package interfaces

import "fmt"

type Teste2 struct {
	S, L string
	N    int
}

func (t Teste2) String() string {
	return fmt.Sprintf("%s, %s", t.S, t.L)
}
