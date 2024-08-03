package main

import (
	"fmt"
)

//	func add(x int, y int) int{
//		return x + y
//	}
func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func calcs(add, sub int) int {
	return add * sub
}

//	func getFullname() (string, string){
//		return "Preet", "Patel"
//	}
func getFullname() (firstName, lastName string) { // specifing the return values' names which can deliver the purpose of those values
	firstName = "Preet"
	lastName = "Patel"
	return // use naked returns only in small functions.
	// return "Preet", "Patel"
}

// func isEligible(age int) string {
// 	// return once only (BAD STYLE)
// 	isEligible := "Nothing"
// 	if age >= 18 && age <= 25 {
// 		isEligible = "You are now eligible for driving!"
// 	} else if age > 25 && age <= 35 {
// 		isEligible = "You are now eligible for weeding!"
// 	} else if age > 35 && age <= 50 {
// 		isEligible = "You are now eligible for kids!"
// 	} else {
// 		isEligible = "You are now not eligible for anything!"
// 	}
// 	return isEligible
// }
func isEligible(age int) string {
	// Use Gardian Cluses
	if age >= 18 && age <= 25 {
		return "You are now eligible for driving!"
	} else if age > 25 && age <= 35 {
		return "You are now eligible for weeding!"
	} else if age > 35 && age <= 50 {
		return "You are now eligible for kids!"
	} else {
		return "You are now not eligible for anything!"
	}

}

func main() {
	fmt.Println("Println is a function in fmt package!")
	// funcitons here mostly works on copy of the values
	fmt.Println(calcs(add(2, 4), sub(4, 1)))

	firstName, _ := getFullname() // avoid the second value by using _

	fmt.Println(firstName)
	fmt.Println(isEligible(20))
}
