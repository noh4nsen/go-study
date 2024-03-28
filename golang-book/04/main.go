package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func add(first int, second int) int {
	return first + second
}

func main() {
	arg := os.Args[1]
	fmt.Printf("%T\n", arg)
	fmt.Println(reflect.TypeOf(arg))

	// String to Integer
	no1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	no2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}

	sum := add(no1, no2)
	fmt.Println(sum)

	fmt.Println(reflect.TypeOf(no1))

	no3String := "100"
	fmt.Println(no3String)

	// Parsing string to different types of integers
	fourBaseEightBitInt, _ := strconv.ParseInt(no3String, 4, 8)
	tenBaseSixteenBitInt, _ := strconv.ParseInt(no3String, 10, 16)
	fmt.Println(reflect.TypeOf(fourBaseEightBitInt))
	fmt.Println(fourBaseEightBitInt)
	fmt.Println(reflect.TypeOf(tenBaseSixteenBitInt))
	fmt.Println(tenBaseSixteenBitInt)

	// Integer to string
	no3 := strconv.Itoa(no2)
	fmt.Println(reflect.TypeOf(no3))
}
