package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Recursive function to generate the sequence
func GenerateSequence(n int, current int) {
	// Base case: Stop if current exceeds n
	if current*current > n {
		return
	}

	// Print the current number or its square
	if current == 2 {
		fmt.Println(current)           // Print the initial value (2)
		fmt.Println(current * current) // Print the square of 2
	} else {
		fmt.Println(current * current) // Print squares of subsequent numbers
	}

	// Recursive call to the next number
	GenerateSequence(n, current+1)
}

func main() {
	var input int
	var err error

	// Check for command-line arguments
	if len(os.Args) > 1 {
		input, err = strconv.Atoi(os.Args[1])
		if err != nil || input <= 0 {
			panic("Please provide a valid positive integer as a command-line argument.")
		}
	} else {
		// Prompt user for input interactively
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a positive integer: ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		// Parse input
		input, err = strconv.Atoi(userInput)
		if err != nil || input <= 0 {
			panic("Please provide a valid positive integer.")
		}
	}

	// Call the recursive function starting with 2
	GenerateSequence(input, 2)
}
