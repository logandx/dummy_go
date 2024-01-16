package main

import (
	"fmt"
	"os"
)

type Bill struct {
	id    int
	tip   float64
	items map[string]float64
	name  string
}

// Constructor func for Bill struct ??
func createBillWithName(name string) Bill {
	user := Bill{
		id:    1,
		tip:   0,
		items: map[string]float64{},
		name:  name,
	}
	return user
}

// / Format fields for Bill struct
func (bill *Bill) format() string {
	formattedStr := "Bill detail: \n"
	var total float64 = 0

	// name
	formattedStr += fmt.Sprintf("%-25v ...%v\n", "Name:", bill.name)
	// loops through the items

	for key, value := range bill.items {
		formattedStr += fmt.Sprintf("%-25v ...$%v\n", key+":", value)
		total += value // sums up the total by values
	}

	// tip
	formattedStr += fmt.Sprintf("%-25v ...$%0.2f \n", "Tip:", bill.tip)

	// total count
	formattedStr += fmt.Sprintf("%-25v ...$%0.2f \n", "Total:", total+bill.tip)
	return formattedStr
}

//* The 2 functions below we could say *(bill).field = value but Go will automatically
//* de-reference it when working with [struct] data type.

// Func to update Bill.name field
func (bill *Bill) updateName(value string) {
	bill.name = value
}

// Func to update the Bill.items field
func (bill *Bill) updateItem(name string, item float64) {
	bill.items[name] = item
}

// Func to update the Bill.tip field
func (bill *Bill) updateTip(value float64) {
	bill.tip = value
}

// Func to save the bill to the specific location
func (bill *Bill) save() {
	data := []byte(bill.format())
	err := os.WriteFile("bills/"+bill.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("The Bill was saved to file")
}
