package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result := sum(1, 2)
	expected := 3
	assert.Equal(t, expected, result)
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
			assert.Equal(t, test.Expected, result)
		})
	}
}
