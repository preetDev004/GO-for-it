package main

import (
	"fmt"
)

// interface is a collection of method signature
type shape interface {
	area() float64
	perimeter() float64
}

// "circle" impliments the shape interface
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

// "rectangle" impliments the shape interface
type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// A function using the interface will work on both (circle and rectangle) as they both sastidfy the shape interface.
func printShapeDetails(s shape) {
	fmt.Printf("Area: %.2f\n", s.area())
	fmt.Printf("Perimeter: %.2f\n", s.perimeter())
}

func main() {
	c := circle{radius: 10}
	r := rectangle{width: 10, height: 20}

	printShapeDetails(c)
	printShapeDetails(r)

	// A type or struct can implement more than 1 interface.

}
