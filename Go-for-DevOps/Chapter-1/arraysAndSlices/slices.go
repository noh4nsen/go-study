package arraysAndSlices

import "fmt"

func Slice(arr []int) {
	printValueAndSize(arr)
	arr = append(arr, 2)
	printValueAndSize(arr)
	arr = append(arr, 2, 3, 4, 6, 7)
	printValueAndSize(arr)
	secondArr := []int{10, 11, 12, 13}
	arr = append(arr, secondArr...)
	printValueAndSize(arr)
	printValueAndSize(arr[2:4])
	printValueAndSize(arr)
	doAppend(arr)
	printValueAndSize(arr)

	for index, val := range arr {
		fmt.Printf("slice entry %d: %d\n", index, val)
	}
}

func printValueAndSize(arr []int) {
	fmt.Println(arr, len(arr))
}

func doAppend(arr []int) {
	arr = append(arr, 100)
	printValueAndSize(arr)
}
