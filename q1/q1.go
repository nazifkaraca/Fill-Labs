package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// countLetterA counts the occurrences of the letter 'a' in a word
func countLetterA(word string) int {
	return strings.Count(strings.ToLower(word), "a")
}

// sortWords sorts words by the count of 'a', length, and lexicographically
func sortWords(words []string) []string {
	sort.SliceStable(words, func(i, j int) bool {
		countA1 := countLetterA(words[i])
		countA2 := countLetterA(words[j])

		// Compare 'a' count
		if countA1 != countA2 {
			return countA1 > countA2
		}

		// Compare length
		if len(words[i]) != len(words[j]) {
			return len(words[i]) > len(words[j])
		}

		// Compare lexicographically
		return words[i] < words[j]
	})
	return words
}

func main() {
	// Prompt user for input
	fmt.Println("Enter words separated by spaces:")

	// Read input from user
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Split the input into words
	words := strings.Fields(input)

	// Sort the words
	sortedWords := sortWords(words)

	// Print the sorted words
	fmt.Println("Sorted words:")
	fmt.Println(sortedWords)
}
