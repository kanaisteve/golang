package main

import (
	"fmt"
	"time"
)

// define a struct
// Employee struct has four fields such as id, name, country and created.
type Employee struct {
	id      int
	name    string
	country string
	created time.Time
}

func main() {
	// declare variable
	var emp Employee
	newEmp := new(Employee)

	// set values
	emp.id = 2
	emp.name = "Kanai"
	emp.country = "Zambia"
	emp.created = time.Now()

	newEmp.id = 5
	newEmp.name = "Kay"
	newEmp.country = "Poland"
	newEmp.created = time.Now()

	// display
	fmt.Println("=========================")
	fmt.Println(emp.id)
	fmt.Println(emp.name)
	fmt.Println(emp.country)
	fmt.Println(emp.created)

	fmt.Println("=========================")
	fmt.Println(newEmp.id)
	fmt.Println(newEmp.name)
	fmt.Println(newEmp.country)
	fmt.Println(newEmp.created)
}

/*
	Structs
	------
	To define a struct, we can use struct keyword.

	type struct_name struct {
		// fields
	}
*/
