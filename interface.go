package main

import (
	"fmt"
	"math"
)

/// Interface example

type Shape interface {
	area() float64
	circumf() float64
}

type Square struct {
	length float64
}

type Circle struct {
	radius float64
}

// Square methods
func (s Square) area() float64 {
	return s.length * s.length
}

func (s Square) circumf() float64 {
	return s.length * 4
}

// Circle methods
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s Shape) {
	fmt.Printf("Area of %T is: %0.2f \n\n", s, s.area())
	fmt.Printf("Circumference of %T is: %0.2f \n\n", s, s.circumf())
}

func createListOfShapes() {
	shapes := []Shape{
		Square{
			length: 4.4,
		},
		Circle{
			4,
		}, Square{
			length: 5,
		},
		Circle{
			8,
		},
	}

	for _, value := range shapes {
		printShapeInfo(value)
	}
}
