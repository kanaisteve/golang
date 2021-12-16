package main

import "fmt"

func main() {
	// iteration - while
	// go doesn't provide while syntax. we can use for
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

/*
	Iteration - while
	------------------
	- Go doesn't provide while syntax like "C family" programming language.
	- We can use for which passes conditional value.
*/
