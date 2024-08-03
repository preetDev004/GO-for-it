package main

import (
	"fmt"
	"os"
	"io"
)

func add(a,b int) int{
	return a+b
}
func mul(a,b int) int{
	return a*b
}

// A functions that accepts a function as a parameter. --> Higher-order function
// A function that has been passed as an argument/data/value to another function. --> First-Class function 
func aggregate(a,b,c int, arithmatic func(int, int) int) int{
	return arithmatic(arithmatic(a,b),c)
}

// A function that takes function as input and returns a new function as output. --> Function Currying
func selfMath(math func(int,int) int) func(int) int{
	return func(x int) int{
		return math(x,x)
	}
}

// Defer keyword allows us to execute some function at the end of the current function or when the current function exists.
// - it is mostly used as a clean up step.
func copyFile(dstName, srcName string) (written int64, err error){
	src, err := os.Open(srcName)
	if err != nil{
		return
	}
	// Close the file when function returns
	defer src.Close() 

	dst, err := os.Create(dstName)
	if err != nil{
		return
	}
	// Close the file when function returns
	defer dst.Close()

	return io.Copy(dst, src)
}

// A function that references variables from outside of its own scope/body. --> Closure function
// - We can store the variables data even after it goes out of scope and use it later on.
// - You can keep track of the count. (Useful for maintaining state)
func adder() func(int) int{
	sum := 0 // outer variable
	return func(x int) int{
		sum += x // available inside the function
		return sum
	}
}

func main(){
	fmt.Println(aggregate(5,2,3,mul))
	fmt.Println(aggregate(5,2,3,add))

	square := selfMath(mul)
	fmt.Println(square(5))

	doubled := selfMath(add)
	fmt.Println(doubled(5))

	w, err := copyFile("./dst.txt", "./src.txt")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(w)

	add := adder()
	fmt.Println(add(5)) 
	// you can keep track and doesnt lose the count.
	fmt.Println(add(1000))
}