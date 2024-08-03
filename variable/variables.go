
package main

import (
	"fmt"
)

func main() {

	// Variables in GO is similar to C/C++.

	var age int = 21 // int is an alias of int32 or int64 depending on the system.
	var t uint16 = 65535 // Unsigned integer of 16 bits. Cannot store value more than 65535.
	fmt.Println(t)

	name := "Preet" // Infer the type automatically upon the value we assign.
	// const pi := 34    // short cut doesn't work for constants.
	const pi float64 = 3.14
	fmt.Println(pi)

	fmt.Println("Hi, I'm ", name)
	// fmt.Printf("My age is %v\n", age) // %v is used for any type.

	msg := fmt.Sprintf("My age is %d", age) // %d is specifically used for int. // Sprintf return the value instead of printing.
	fmt.Println(msg)

}