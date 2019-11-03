package main

import (
	"fmt"
	"github.com/material-seminario-golang/fundamentals/functions"
	"github.com/material-seminario-golang/fundamentals/loops"
	"github.com/material-seminario-golang/fundamentals/pointers"
	"github.com/material-seminario-golang/fundamentals/structs"
)

func main() {
	fmt.Println("hello fundamentals")

	// Functions
	fmt.Println("explicit", functions.SquaredExplicit(4))
	fmt.Println("implicit", functions.SquaredImplicit(5))

	// Loops and Arrays - Beware of last loop!!
	loops.PrintLoops()

	// Structures
	car := structs.NewCar("Fiat", "147", 150)
	car.Accelerate()
	car.Accelerate()
	car.Brake()
	fmt.Println(car.Speed)

	// Pointers!
	pointers.WorkWithPointers()
}
