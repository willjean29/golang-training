package main

import "fmt"

func mains() {
	prices := []float64{10.99, 9.99}
	prices[1] = 11.99
	// dispatch erro
	// prices[2] = 12.99
	prices = append(prices, 12.99)
	fmt.Println(prices)

	discountPrices := []float64{20, 12, 34}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// slices with static sizes
// func main() {
// 	var productNames [4]string = [4]string{"A Books"}
// 	prices := [4]float64{1.99, 2.99, 3.99, 4.99}
// 	productNames[2] = "A Carpet"
// 	fmt.Println(prices[2])
// 	fmt.Println(productNames)

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))

// }
