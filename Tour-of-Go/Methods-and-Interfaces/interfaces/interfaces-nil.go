package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
	var t *T

	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	// Calling a method on a nil interface is a runtime error
	// because there is no type inside the interface tuple to
	// indicate which concrete method to call
	//var u I
	//describe(u)
	//u.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
