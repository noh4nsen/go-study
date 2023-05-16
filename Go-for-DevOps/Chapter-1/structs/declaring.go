package structs

import "fmt"

func PrintStruct() {
	var record = struct {
		Name string
		Age  int
	}{
		Name: "John Doak",
		Age:  100,
	}

	fmt.Printf("%s is %d years old\n", record.Name, record.Age)
}
