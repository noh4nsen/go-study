package main

import "fmt"

func main() {
	var response1, response2 string

	fmt.Scan(&response1, &response2)
	str := fmt.Sprintf("response1: %s ; response2: %s", response1, response2)
	fmt.Println(str)

	var prefix string
	var no int

	fmt.Scanf("%3s%d", &prefix, &no)
	fmt.Printf("prefix: %s, invoice no: %d\n", prefix, no)
}
