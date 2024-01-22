package main

import "fmt"

func main() {
	age := 32
	var agePointer *int
	agePointer = &age
	fmt.Println("Age: ", *agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println("Adult years: ", adultYears)

	fmt.Println("Age: ", *agePointer)

}

func getAdultYears(age *int) int {
	*age = *age - 18
	return *age
}
