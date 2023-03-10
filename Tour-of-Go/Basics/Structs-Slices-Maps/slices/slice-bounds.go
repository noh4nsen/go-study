package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	t := s[1:4]
	fmt.Println(t)

	u := s[:2]
	fmt.Println(u)

	v := s[1:]
	fmt.Println(v)
}
