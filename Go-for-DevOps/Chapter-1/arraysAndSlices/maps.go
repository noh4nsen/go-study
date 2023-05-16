package arraysAndSlices

import "fmt"

func Maps() {
	counters := make(map[string]int, 10)

	modelToMake := map[string]string{
		"prius":    "toyota",
		"chevelle": "chevy",
	}

	fmt.Println(counters)
	fmt.Println(modelToMake)

	fmt.Println(modelToMake["chevelle"])
	fmt.Println(modelToMake["outback"])
	fmt.Println(modelToMake["prius"])

	modelToMake["outback"] = "subaru"
	counters["pageHits"] = 10
	fmt.Println(modelToMake)
	fmt.Println(counters)

	for key, val := range modelToMake {
		fmt.Printf("car model %q has make %q\n", key, val)
	}
}
