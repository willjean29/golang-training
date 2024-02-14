package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Printf("Hello, world.\n")
}

func sum(a, b int) int {
	return a + b
}

func sumStr(a, b string) int {
	time.Sleep(3 * time.Second)
	c, err := strconv.Atoi(a)
	if err != nil {
		return 0
	}
	d, err := strconv.Atoi(b)
	if err != nil {
		return 0
	}
	return c + d
}

type Square struct {
	length int
}

func (s *Square) Area() int {
	return s.length * s.length
}

func (s *Square) Perimeter() int {
	return 4 * s.length
}
