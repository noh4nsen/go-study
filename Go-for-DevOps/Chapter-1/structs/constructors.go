package structs

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) (*Record, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be the empty string")
	}
	if age <= 0 {
		return nil, fmt.Errorf("age cannot be less than zero")
	}

	return &Record{Name: name, Age: age}, nil
}
