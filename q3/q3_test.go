package main

import (
	"testing"
)

func TestMostRepeated(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "Test Case 1",
			input:    []string{"apple", "pie", "apple", "red", "red", "red"},
			expected: "red",
		},
		{
			name:     "Test Case 2",
			input:    []string{"blue", "blue", "green", "green", "blue"},
			expected: "blue",
		},
		{
			name:     "Test Case 3",
			input:    []string{"cat", "dog", "cat", "cat", "dog", "dog", "dog"},
			expected: "dog",
		},
		{
			name:     "Test Case 4",
			input:    []string{"single"},
			expected: "single",
		},
		{
			name:     "Test Case 5",
			input:    []string{},
			expected: "", // Edge case: empty slice
		},
	}

	for _, test := range tests {
		result := mostRepeated(test.input)

		if result != test.expected {
			t.Errorf("%s: FAILED. Expected %v, Got %v", test.name, test.expected, result)
		} else {
			t.Logf("%s: PASSED", test.name)
		}
	}
}
