package employee_test

import (
	"app/testing/employee"
	"testing"
)

// func TestValidate_EmptyName(t *testing.T) {
// 	e := employee.Employee{
// 		Name: "",
// 	}

// 	isValid, err := employee.Validate(e)

// 	if isValid != false {
// 		t.Error("Expected isValid to be false, got true")
// 	}

// 	if err == nil {
// 		t.Error("Expected error to be not nil, got nil")
// 	}

// 	if err.Error() != "name cannot be empty" {
// 		t.Errorf("Expected err.Error() to be 'name cannot be empty', got %s", err.Error())
// 	}
// }

// func TestValidate_WrongEmail(t *testing.T) {
// 	e := employee.Employee{
// 		Name:  "Jhon Doe",
// 		Email: "",
// 	}

// 	isValid, err := employee.Validate(e)

// 	if isValid != false {
// 		t.Error("Expected isValid to be false, got true")
// 	}

// 	if err == nil {
// 		t.Error("Expected error to be not nil, got nil")
// 	}

// 	if err.Error() != "email must contain @" {
// 		t.Errorf("Expected err.Error() to be 'email must contain @', got %s", err.Error())
// 	}
// }

// func TestValidate_WrongAge(t *testing.T) {
// 	e := employee.Employee{
// 		Name:  "John Doe",
// 		Age:   15,
// 		Email: "jhon@gmail.com",
// 	}

// 	isValid, err := employee.Validate(e)

// 	if isValid != false {
// 		t.Error("Expected isValid to be false, got true")
// 	}

// 	if err == nil {
// 		t.Error("Expected error to be not nil, got nil")
// 	}

// 	if err.Error() != "age must be greater than 18" {
// 		t.Errorf("Expected err.Error() to be 'age must be greater than 18', got %s", err.Error())
// 	}
// }

// func TestValidate(t *testing.T) {
// 	e := employee.Employee{
// 		Name:  "John Doe",
// 		Age:   25,
// 		Email: "jhon@gmail.com",
// 	}

// 	isValid, err := employee.Validate(e)

// 	if isValid != true {
// 		t.Error("Expected isValid to be true, got false")
// 	}

// 	if err != nil {
// 		t.Error("Expected error to be nil, got error")
// 	}
// }

func TestValidateTotal(t *testing.T) {
	tests := []struct {
		name     string
		employee employee.Employee
		wantErr  bool
	}{
		{"Valid Employee", employee.Employee{"John Doe", 19, "john.doe@example.com"}, false},
		{"Empty Name", employee.Employee{"", 19, "john.doe@example.com"}, true},
		{"Invalid Email", employee.Employee{"John Doe", 19, "johndoeexample.com"}, true},
		{"Invalid Age", employee.Employee{"John Doe", 17, "john.doe@example.com"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid, err := employee.Validate(tt.employee)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if isValid != !tt.wantErr {
				t.Errorf("Validate() = %v, want %v", isValid, !tt.wantErr)
			}
		})
	}
}
