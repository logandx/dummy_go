package main

import "fmt"

type Bill struct {
	id    int
	item  float64
	items map[string]float64
	name  string
}

// Constructor func for Bill struct ??
func createUserWithName(name string) Bill {
	user := Bill{
		id:    1,
		item:  0,
		items: map[string]float64{},
		name:  name,
	}
	return user
}

// / Format fields for Bill struct
func (bill *Bill) format() string {
	formattedStr := "Bill detail: \n"
	var total float64 = 0

	// loops through the items
	for key, value := range bill.items {
		formattedStr += fmt.Sprintf("%-25v ...$%v\n", key+":", value)
		total += value // sums up the total by values
	}

	// name
	formattedStr += fmt.Sprintf("%-25v ...%v\n", "Name:", bill.name)

	// total count
	formattedStr += fmt.Sprintf("%-25v ...$%0.2f \n", "Total:", total+bill.item)
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
