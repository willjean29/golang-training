package main

import (
	"bufio"
	"fmt"
	"interfaces/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

func main() {
	title := getTodoData()
	userTodo, err := todo.New(title)
	if err != nil {
		fmt.Println(err)
		return
	}
	userTodo.Display()
	// err = userTodo.Save()
	// if err != nil {
	// 	fmt.Println("Saving the todo failed")
	// 	return
	// }
	// fmt.Println("Saving the todo succeeded!")
	err = saveData(userTodo)
	if err != nil {
		return
	}
}

func getTodoData() string {
	title := getUserInput("Todo title:")
	return title
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}
	fmt.Println("Saved the note successfully")
	return nil
}
