package main

import (
	"errors"
	"fmt"
)

// Define a custom type for the allowed constant strings
type AllowedString string

const (
	FirstString  AllowedString = "first string"
	SecondString AllowedString = "second string"
	ThirdString  AllowedString = "third string"
	FourthString AllowedString = "fourth string"
)

func validateInput(input AllowedString) error {
	// Validate the input against the allowed constant strings
	switch input {
	case FirstString, SecondString, ThirdString, FourthString:
		return nil // Input is valid
	default:
		return errors.New("input is not one of the allowed constant strings")
	}
}

func main() {
	// Valid input
	if err := validateInput(ThirdString); err == nil {
		fmt.Println("Input is valid")
	} else {
		fmt.Println("Error:", err)
	}

	// Invalid input
	if err := validateInput("fifth string"); err == nil {
		fmt.Println("Input is valid")
	} else {
		fmt.Println("Error:", err)
	}
}
