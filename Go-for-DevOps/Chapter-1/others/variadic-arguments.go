package others

import "fmt"

func Variadic() {
	args := []int{5, 4, 3, 2, 1}
	res := sum(args)
	fmt.Printf("Result: %d", res)

	fmt.Println()

	res = sumVariadic(5, 4, 3, 2, 1)
	fmt.Printf("Result: %d", res)
}

func sum(numbers []int) int {
	sum := 0

	for i, n := range numbers {
		sum += n
		fmt.Printf("Index: %d, Value: %d, Sum: %d\n", i, n, sum)
	}

	return sum
}

func sumVariadic(numbers ...int) int {
	sum := 0

	for i, n := range numbers {
		sum += n
		fmt.Printf("Index: %d, Value: %d, Sum: %d\n", i, n, sum)
	}

	return sum
}
