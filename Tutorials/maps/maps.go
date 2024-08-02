package main 

import (
	"fmt"
)

func main(){

	// Creating maps
	// marks := make(map[string]int)
	// marks["preet"] = 95
	// marks["bachu"] = 90
	// marks["maddy"] = 99

	marks := map[string]int{
		"preet": 95,
		"bachu": 90,
		"maddy": 99,
	}
	fmt.Println(marks, "-> length of map: ", len(marks))

	// deleting an element
	delete(marks, "maddy")
	fmt.Println(marks)

	// check if the element exists
	if elem, ok := marks["preet"]; ok{
		fmt.Println("Element exists, value -> ", elem)
	}else{
		fmt.Println("Element doesn't exist")
	}

	// nested maps
	students := map[string]map[string]int{
		"preet": {
			"marks": 95,
			"age": 21,
		},
		"bachu": {
			"marks": 90,
			"age": 22,
		},
	}
	fmt.Println(students)
	fmt.Println(students["preet"]["marks"])

	// rather then using nested maps use Structs as the type
	type student struct{
		marks ,age int
	}
	pupils := map[string]student{
		"preet": {
			marks: 95,
			age: 21,
		},
		"bachu": student{ // it's okay if you dont specify student type
			marks: 90,
			age: 22,
		},
	}
	// Code will not panic if we try to access a key which is not present in the map. 
	// (returns 0 value of the type)
	fmt.Println(pupils)
	fmt.Println(pupils["preet"].age)

	// we use rune rather than having a string of 1 length.
	char := 'a'
	// var char rune = '3' 
	fmt.Println(char)

}
// NOTE : Unlike other programming languages you dont need to pass mpas as pointers to other func. 