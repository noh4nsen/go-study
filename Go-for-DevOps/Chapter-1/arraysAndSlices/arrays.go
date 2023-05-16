package arraysAndSlices

import "fmt"

func ChangeValueAtZeroIndex(array [2]int) {
	array[0] = 3

	fmt.Println("inside: ", array[0])
}
