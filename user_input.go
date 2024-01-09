package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/// Bill input

func createUser() Bill {
	// Create new reader from standard input package
	reader := bufio.NewReader(os.Stdin)

	input, err := getUserInput("Create new user with name - Name: ", reader)
	if err != nil {
		return Bill{}
	}

	user := createUserWithName(input)
	fmt.Println("New bill created - Name: ", user.name)
	return user
}

// Helper func to get the user input
func getUserInput(prompt string, newReader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := newReader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input, err
}

// Func to choose the prompt options by the user input
func promptOptions(user Bill) {
	reader := bufio.NewReader(os.Stdin)

	options, err := getUserInput("Choose an option (a - Add Name, s - Save the Bill, t - Add Tip): ", reader)
	if err != nil {
		return
	}
	switch options {
	case "a":
		name, _ := getUserInput("Type Item name:", reader)
		price, _ := getUserInput("Type Item price:", reader)
		fmt.Println(name, price)
	case "t":
		fmt.Println("You chose 't'")
	case "s":
		tip, _ := getUserInput("Enter tip amount ($): ", reader)
		fmt.Println(tip)
	default:
		fmt.Println("Invalid option...")
		promptOptions(user)
	}

	fmt.Println(options)
}
