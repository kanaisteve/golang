package main

import "fmt"

func main() {
	// iteration - for
	var i int
	for i = 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("First iteration ending...")

	for j := 5; j < 11; j++ {
		fmt.Println(j)
	}
}

/*
	Iterationn - for
	-----------------
	Iteration operation is useful when we do repetitive activities.
	The following is for syntax:

	for initialization;conditional;increment/decrement {
		...
	}
*/
