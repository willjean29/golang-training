package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		go getData(i)
	}
	var s string
	fmt.Scanln(&s)
	duration := time.Since(start)
	fmt.Println("Duration: ", duration)

}

func hi(num int) {
	fmt.Println("Hello ", num)
	time.Sleep(1 * time.Second)
}

func getData(id int) {
	fmt.Println("Getting data from user ", id)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/" + strconv.Itoa(id))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}
