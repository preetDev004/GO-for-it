package main

import (
	"fmt"
)

func modifySlice(slice []int) {
	slice[0] = 10
}

// Normaly used to work with slices. Flexible for caller.
func variadicModifySlice(slice ...int) {
	slice[1] = 50
}

func main() {

	// In GO arrays are fixed size.
	arr := [6]int{1, 3, 4, 5, 3, 2}
	fmt.Println("Length of Array: ", len(arr))

	// slice := []int{1, 3, 4, 5, 3, 4, 2} // This is slices in GO. They don't have fix size.
	slice := arr[2:4] // This is a slice built on top of array.
	fmt.Println("Length of Slice: ", len(slice))

	// converting an array into slice
	sliceFromArray := arr[:]
	fmt.Println("Length of Slice from Array: ", len(sliceFromArray))

	// NOTE : Slices are paas by ref so if you pass it to func and if it changes something in it, it will be reflected in the caller func as well.
	fmt.Println("Array before func call: ", arr)
	fmt.Println("Slice before func call: ", slice)

	modifySlice(slice)
	// modifySlice(1,3,4) you can't do this!!
	variadicModifySlice(slice...) // "..." spread operator to open the slice and pass it to the func.
	variadicModifySlice(1, 3, 4)

	fmt.Println("Slice after func call: ", slice)
	fmt.Println("Array after func call: ", arr)

	// Creates a slice with length 5. Here, underlying array has the capacity of 10.
	mySlice := make([]int, 5, 10) // don't have to create the array of size 10 explicitly. make() will do it.
	// mySlice := make([]int ,5) // if you don't specify the capacity it will be set to the length.
	fmt.Println("Length of mySlice: ", len(mySlice))
	fmt.Println("Capacity of mySlice: ", cap(mySlice))

	// Appending - Always assign the appended slice to the original ones!
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 8, 5)
	mySlice = append(mySlice, slice...)

	fmt.Println("After appending: ", mySlice)

	// for i := 0; i<len(slice);i++{
	// 	fmt.Println(slice[i])
	// }
	for i, num := range slice { // iterate over the slice with range keyword.
		fmt.Println(i, num)
	}
}
