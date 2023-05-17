package structs

import "fmt"

type CarModel string

func SimpleType() {
	var myCar CarModel = "Chevelle"
	fmt.Println(myCar)

	myCar = CarModel("Chevette")
	fmt.Println(myCar)

	myCarAsString := string(myCar)
	fmt.Println(myCarAsString)
}

type Record struct {
	Name string
	Age  int
}

func (r *Record) String() string {
	return fmt.Sprintf("%s, %d", r.Name, r.Age)
}

func (r *Record) IncrAge() {
	r.Age++
}

func StructType() {
	david := Record{Name: "David Justice", Age: 28}
	sarah := Record{Name: "Sarah Murphy", Age: 29}

	fmt.Printf("%+v\n", david)
	fmt.Printf("%+v\n", sarah)

	fmt.Println(david.String())
	fmt.Println(sarah.String())

	david.Name = "Testing value change"
	fmt.Println(david.String())

	david.IncrAge()
	david.IncrAge()
	fmt.Println(david.String())

	PrintStruct()
}

func ChangeName(r *Record) {
	r.Name = "Peter"
	fmt.Println("inside ChangeName: ", r.Name)
}
