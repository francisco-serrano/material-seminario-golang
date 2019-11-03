package pointers

import "fmt"

func DoSomeWorkNoPointers(value int) {
	fmt.Println("DoSomeWorkNoPointers", value)
	value = 6
	fmt.Println("DoSomeWorkNoPointers", value)
}

func DoSomeWorkPointers(value *int) {
	fmt.Println("DoSomeWorkPointers", *value)
	*value = 6
	fmt.Println("DoSomeWorkPointers", *value)
}

func WorkWithPointers() {
	var value int
	value = 5

	fmt.Println("WorkWithPointers", value)
	DoSomeWorkPointers(&value)
	fmt.Println("WorkWithPointers", value)

	fmt.Println("-----------------------")

	value = 5

	fmt.Println("WorkWithPointers", value)
	DoSomeWorkNoPointers(value)
	fmt.Println("WorkWithPointers", value)
}
