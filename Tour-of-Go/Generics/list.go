package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	a := List[int]{nil, 2}
	b := List[int]{nil, 3}
	c := List[int]{nil, 4}
	a.next = &b
	b.next = &c
	c.next = &a

	fmt.Printf("A ::: Mem :: %v; Val :: %v; Next :: %v\n", &a, a.val, a.next)
	fmt.Printf("B ::: Mem :: %v; Val :: %v; Next :: %v\n", &b, b.val, b.next)
	fmt.Printf("C ::: Mem :: %v; Val :: %v; Next :: %v\n", &c, c.val, c.next)
}
