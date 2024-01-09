package main

import "fmt"

/// Maps

var menu map[string]float64 = map[string]float64{
	"soup":   4.99,
	"pie":    6.11,
	"salad":  5.5,
	"coffee": 7.2,
}

func mapTest() {

	fmt.Println(menu)

}
