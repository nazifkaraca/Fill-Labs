package main

import "fmt"

// mostRepeated finds the most repeated element in a given string slice
func mostRepeated(data []string) string {
	// Map to store frequency of each element
	frequency := make(map[string]int)

	// Variable to track the most repeated element and its count
	var mostRepeatedElement string
	maxCount := 0

	// Count occurrences of each element
	for _, item := range data {
		frequency[item]++
		if frequency[item] > maxCount {
			maxCount = frequency[item]
			mostRepeatedElement = item
		}
	}

	return mostRepeatedElement
}

func main() {
	// Modify here when needed
	data := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	result := mostRepeated(data)
	fmt.Printf("The most repeated element is: %s\n", result)
}
