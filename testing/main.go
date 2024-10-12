package main

import (
	"fmt"

	"app/testing/employee"
)

func main() {
	ep := employee.Employee{
		Name:  "John Doe",
		Age:   15,
		Email: "jhon@gmail.com",
	}

	isValid, err := employee.Validate(ep)
	fmt.Println()
	fmt.Println("Is valid: ", isValid)
	fmt.Println("Error: ", err.Error())
}
