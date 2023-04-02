package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a reader to read from the standard input
	reader := bufio.NewReader(os.Stdin)

	// Print a welcome message
	fmt.Println("Hello, this is a simple CLI program in Golang.")
	fmt.Println("Type 'help' to see the available commands.")
	fmt.Println("Type 'exit' to quit the program.")

	// Loop until the user types 'exit'
	for {
		// Prompt the user to enter a command
		fmt.Print("Enter command: ")

		// Read the user input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Trim the input and convert it to lower case
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)

		// Call the handleCommand function with the input and get the return value
		done := handleCommand(input)

		// If the return value is true, exit the loop
		if done {
			break
		}
	}
}

// Define the handleCommand function that takes a string argument and returns a bool value
func handleCommand(input string) bool {
	// Handle the input based on the command

	result := false

	switch input {
	case "help":
		// Print the help message
		fmt.Println("Available commands:")
		fmt.Println("- help: show this message")
		fmt.Println("- hello: greet the user")
		fmt.Println("- exit: quit the program")
		// Return false to indicate the program is not done
		result = false
	case "hello":
		// Print a greeting message
		fmt.Println("Hello, nice to meet you!")
		// Return false to indicate the program is not done
		result = false
	case "exit":
		// Print a farewell message
		fmt.Println("Bye, have a good one!")
		// Return true to indicate the program is done
		result = true
	default:
		// Print an error message for unknown commands
		fmt.Println("Invalid command:", input)
		fmt.Println("Type 'help' to see the available commands.")
		// Return false to indicate the program is not done
		result = false
	}
	return result
}
