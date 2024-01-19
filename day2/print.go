package main

import "fmt"

func main() {
	var name, lastName string = "Jean", "Osco"
	edad := 27
	// output text with formatting
	// fmt.Sprintf ==> prepare output text and save in variable
	// fmt.Printf ==> prepare output text and show on terminal
	// `string literal` => take format string
	message := fmt.Sprintf("Hola %s %s, tienes %d años de edad\n", name, lastName, edad)
	fmt.Println(message)
	outputText("Hello, World!")
	sum, sub := calculate(10, 5)
	fmt.Println(sum, sub)
}

func outputText(message string) {
	fmt.Println(message)
}

func calculate(a, b int) (sum int, sub int) {
	// return a + b, a - b
	sum = a + b
	sub = a - b
	return
}
