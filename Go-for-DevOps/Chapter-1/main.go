package main

import (
	"Chapter-1/arraysAndSlices"
	"Chapter-1/others"
	"Chapter-1/pointers"
	"Chapter-1/say"
	"Chapter-1/structs"
	"fmt"
)

func main() {
	fmt.Println("::: Others :::")
	fmt.Println("\n- Functions")
	others.Functions()
	fmt.Println("\n- Anonymouse Functions")
	others.AnonymousFunctions()
	fmt.Println("\n- Variable Shadowing")
	others.VariableShadowing()
	fmt.Println("\n- Variadics")
	others.Variadic()

	fmt.Println("\n\n::: Say (private and public functions) :::")
	say.PrintHello()
	say.PrintHelloWorld()

	fmt.Println("\n\n::: Arrays And Slices :::")
	fmt.Println("\n- Arrays")
	arraysAndSlices.ChangeValueAtZeroIndex([2]int{})
	fmt.Println("\n- Slices")
	arraysAndSlices.Slice([]int{})
	fmt.Println("\n- Maps")
	arraysAndSlices.Maps()

	fmt.Println("\n\n::: Pointers :::")
	fmt.Println("\n- Memory Addresses")
	pointers.MemoryAddress(3)
	word := "hello "
	pointers.ChangeValue(&word)
	fmt.Println("\n", word)

	fmt.Println("\n\n::: Structs :::")
	structs.PrintStruct()
	structs.SimpleType()
	structs.StructType()
}
