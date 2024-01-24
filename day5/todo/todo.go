package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Title string `json:"title"`
}

func (todo Todo) Display() {
	fmt.Printf("Your note titled %v \n", todo.Title)
}

func (todo Todo) Save() error {
	filename := strings.ReplaceAll(todo.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)
}

func New(title string) (Todo, error) {
	if title == "" {
		return Todo{}, errors.New("Invalid note data.")
	}
	return Todo{
		Title: title,
	}, nil
}
