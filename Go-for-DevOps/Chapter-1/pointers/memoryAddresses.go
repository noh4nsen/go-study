package pointers

import "fmt"

func MemoryAddress(x int) {
	fmt.Printf("Memory Address = %p; Value = %d", &x, x)
}
