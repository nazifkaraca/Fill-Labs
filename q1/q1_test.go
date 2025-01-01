package main

import (
	"reflect"
	"strings"
	"testing"
)

// Lower input for comparison
func normalizeToLower(words []string) []string {
	lowercaseWords := make([]string, len(words))
	for i, word := range words {
		lowercaseWords[i] = strings.ToLower(word)
	}
	return lowercaseWords
}

// Test cases written in name, input and expected format
func TestSortWords(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Test Case 1",
			input:    []string{"aaaasd", "a", "aab", "aaabcd", "ef", "csssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"},
			expected: []string{"aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "csssssssd", "fdz", "ef", "kf", "zc", "l"},
		},
		{
			name:     "Test Case 2",
			input:    []string{"abc", "aaaa", "a", "aaa", "aa", "b"},
			expected: []string{"aaaa", "aaa", "aa", "abc", "a", "b"},
		},
		{
			name:     "Test Case 3",
			input:    []string{"banana", "apple", "grape", "aaa", "aa", "a"},
			expected: []string{"banana", "aaa", "aa", "apple", "grape", "a"},
		},
		{
			name:     "Test Case 4",
			input:    []string{"aaaa"},
			expected: []string{"aaaa"},
		},
		{
			name:     "Test Case 5",
			input:    []string{"xyz", "bcd", "def", "klmn", "pqrst"},
			expected: []string{"pqrst", "klmn", "bcd", "def", "xyz"}, // Sorted by length in descending order
		},
		{
			name:     "Test Case 6",
			input:    []string{"a", "ab", "abc", "abcd", "abcde"},
			expected: []string{"abcde", "abcd", "abc", "ab", "a"}, // Sorted by length in descending order
		},
		{
			name:     "Test Case 7",
			input:    []string{"aaa", "aa", "banana", "a", "longword", "abcd", "aab"},
			expected: []string{"banana", "aaa", "aab", "aa", "abcd", "a", "longword"},
		},
		{
			name:     "Test Case 8",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "Test Case 9",
			input:    []string{"a!a!a", "aaa", "a a a", "a", "a-aa", "ab", "xyz"},
			expected: []string{"a a a", "a!a!a", "a-aa", "aaa", "ab", "a", "xyz"},
		},
		{
			name:     "Test Case 10",
			input:    []string{"AAA", "AaA", "aaa", "A", "a"},
			expected: []string{"aaa", "AaA", "AAA", "a", "A"}, // Assuming case insensitivity
		},
		{
			name:     "Test Case 11",
			input:    []string{"a", "b", "c", "a", "b", "a"},
			expected: []string{"a", "a", "a", "b", "b", "c"}, // 'a' comes first, then the others
		},
		{
			name:     "Test Case 12",
			input:    []string{"aaa", "aaa", "aa", "aa", "a", "a"},
			expected: []string{"aaa", "aaa", "aa", "aa", "a", "a"},
		},
		{
			name:     "Test Case 13",
			input:    []string{"aabb", "aabc", "aaab", "abcd"},
			expected: []string{"aaab", "aabb", "aabc", "abcd"}, // Sorted by 'a' count
		},
	}

	// Implement test cases
	for _, test := range tests {
		normalizedInput := normalizeToLower(test.input)       // Normalize input to lower
		normalizedExpected := normalizeToLower(test.expected) // Normalize expected to lower
		result := sortWords(normalizedInput)                  // Result

		// Test whether input and expected equal
		if !reflect.DeepEqual(result, normalizedExpected) {
			t.Errorf("%s: FAILED. Expected %v, Got %v", test.name, normalizedExpected, result)
		} else {
			t.Logf("%s: PASSED", test.name)
		}
	}
}
