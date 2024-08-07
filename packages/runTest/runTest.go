package main

import (
	"fmt"
	"github.com/preetDev004/GO-for-it/test"

)
func printName(){
	fmt.Println("Hello, I am Preet!")
}

func main(){
	test.Run(printName)
}