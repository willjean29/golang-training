package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	// one form of declaration variable
	var investmentAmount, years float64
	// another form of declaration and assignment variable
	expectedReturnRate := 5.5
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	// validate type of data
	// if(reflect.TypeOf(investmentAmount).Kind() == reflect.Float64) {
	// 	fmt.Println("Investment Amount is a float required")
	// 	return
	// }
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+ expectedReturnRate/100),years)
	futureRealValue := futureValue / math.Pow((1+ inflationRate/100),years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}