package employee

import (
	"fmt"
	"strings"
)

type Employee struct {
	Name  string
	Age   int
	Email string
}

func Validate(e Employee) (isValid bool, err error) {
	if e.Name == "" {
		return false, fmt.Errorf("name cannot be empty")
	}

	if !strings.Contains(e.Email, "@") {
		return false, fmt.Errorf("email must contain @")
	}

	if e.Age < 18 {
		return false, fmt.Errorf("age must be greater than 18")
	}

	return true, nil
}
