package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	//mapTest()
	//firstName1, secondName1 := getInitials("asd sdf")
	//fmt.Println("First capital letter of names:", firstName1, secondName1)
	//
	//firstName2, secondName2 := getInitials("asd")
	//fmt.Println("First capital letter of names:", firstName2, secondName2)

	/// Struct example

	//user1 := createUserWithName("Oga")
	//
	//user1.updateName("Tatsumi Oga")
	//user1.updateItem("Lol", 5.5)
	//user1.updateItem("Lmao", 8.5)
	//user1.updateItem("Hihi", 5.6)
	//user1.updateItem("Wtf", 1.5)
	//user1.updateItem("?", 2.5)
	//
	//fmt.Printf(user1.format())

	/// Bill input
	newUser := createUser()
	promptOptions(newUser)
	fmt.Println(newUser)

}

//
//func sayGreeting(name string) {
//	fmt.Printf("Good moring %v \n", name)
//}
//
//func cycleNames(names []string, f func(value string)) {
//	for _, value := range names {
//		f(value)
//
//	}
//}
//
//func circleArea(radius float64) float64 {
//	area := math.Pi * radius * radius
//	return area
//}
