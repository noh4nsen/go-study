package others

import "fmt"

func Functions() {
	result := add(3, 2, "teste")
	fmt.Println(result)

	res, rem := divide(10, 3)
	fmt.Printf("Result: %d, Remainder: %d\n", res, rem)
}

func add(x, y int, s string) int {
	fmt.Println(s)
	return x + y
}

func divide(num, div int) (res, rem int) {
	res = num / div
	rem = num % div
	return res, rem
}
