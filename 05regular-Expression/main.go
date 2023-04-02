package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Define an array of emails to check
	emails := []string{
		"alice@example.com",
		"bob@gmail",
		"charlie@domain.co.uk",
		"david@invalid.",
		"eve@sub.domain.com",
		"frank@.com",
	}

	// Loop through the emails and print the validation result
	for _, email := range emails {
		fmt.Printf("%s is valid: %t\n", email, isValidEmail(email))
	}
}

// Define the isValidEmail function that takes a string argument and returns a bool value
func isValidEmail(email string) bool {
	// Define a regular expression to match a valid email format
	// Source: https://www.geeksforgeeks.org/check-if-an-email-address-is-valid-or-not-in-golang/
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	// Return the result of matching the email with the regular expression
	return regex.MatchString(email)
}
