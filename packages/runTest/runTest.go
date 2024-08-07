package main

import (
	"fmt"
	"github.com/preetDev004/GO-for-it/run"

)
func printName(){
	fmt.Println("Hello, I am Preet!")
}

func main(){
	run.Run(printName)
}