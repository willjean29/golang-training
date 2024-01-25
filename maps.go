package main

import "fmt"

// uso de alias
type floatMap map[string]float64

func (f floatMap) output() {
	fmt.Println(f)
}

func main() {
	userNames := make([]string, 2, 5)
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")
	fmt.Println(userNames)
	coursesRating := make(floatMap, 3)
	coursesRating["Go"] = 5
	coursesRating["Python"] = 4
	coursesRating["Java"] = 3
	coursesRating.output()

	for _, user := range userNames {
		fmt.Println(user)
	}

	for key, value := range coursesRating {
		fmt.Println(key, value)
	}
}

// usign map and added, deleted data
// func main() {
// 	websites := map[string]string{
// 		"Google":   "http://google.com",
// 		"Facebook": "http://facebook.com",
// 	}

// 	fmt.Println(websites)
// 	fmt.Println(websites["Google"])

// 	websites["Amazon Web Services"] = "http://aws.com"

// 	fmt.Println(websites)

// 	delete(websites, "Facebook")

// 	fmt.Println(websites)
// }
