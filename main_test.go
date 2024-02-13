package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(1, 2)
	if result != 3 {
		t.Errorf("sum was incorrect, got: %v want: %v ", result, 3)
	}
}

func TestSumStr(t *testing.T) {
	testCases := []struct {
		Name     string
		A        string
		B        string
		Expected int
	}{
		{"Valid", "1", "2", 3},
		{"Invalid A", "a", "2", 0},
		{"Invalid B", "1", "b", 0},
	}

	for i := range testCases {
		test := testCases[i]
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			result := sumStr(test.A, test.B)
			if result != test.Expected {
				t.Errorf("sumStr was incorrect, got: %v want: %v ", result, test.Expected)
			}
		})
	}
}
