package loops

import "fmt"

func PrintLoops() {
	values := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(values); i++ {
		fmt.Println(values[i])
	}

	var i int = 0
	for i < len(values) {
		fmt.Println(values[i])

		i++
	}

	for {
		fmt.Println(values)
	}
}
