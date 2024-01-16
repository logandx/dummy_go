package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/// Bill input

func createBill() Bill {
	// Create new reader from standard input package
	reader := bufio.NewReader(os.Stdin)

	input, err := getUserInput("Create new bill with name - Name: ", reader)
	if err != nil {
		return Bill{}
	}

	bill := createBillWithName(input)
	fmt.Println("New bill created - Name: ", bill.name)
	return bill
}

// Helper func to get the user input
func getUserInput(prompt string, newReader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := newReader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input, err
}

// Func to choose the prompt options by the user input
func promptOptions(bill Bill) {
	reader := bufio.NewReader(os.Stdin)

	options, err := getUserInput("Choose an option (a - Add Name, s - Save the Bill, t - Add Tip): ", reader)
	if err != nil {
		return
	}
	switch options {
	case "a":
		name, _ := getUserInput("Type Item name:", reader)
		price, _ := getUserInput("Type Item price:", reader)
		parsedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(bill)
		}
		bill.updateItem(name, parsedPrice)
		fmt.Println("Item added - ", name, price)
		promptOptions(bill)
	case "t":
		tip, _ := getUserInput("Enter tip amount ($): ", reader)
		parsedTip, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(bill)
		}
		bill.updateTip(parsedTip)
		fmt.Println("Tip added - ", tip)
		promptOptions(bill)
	case "s":
		bill.save()
		fmt.Println("The Bill has been saved - Name: ", bill.name)
	default:
		fmt.Println("Invalid option...")
		promptOptions(bill)
	}

	fmt.Println(options)
}
