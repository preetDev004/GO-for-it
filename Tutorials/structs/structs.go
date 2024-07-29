package main

import (
	"fmt"
)

type car struct{
	brand string
	model string
	year int
	price float64
}

// embded structs
type truck struct{
	car
	capacity float64
}

func (c car) displayDetails(){
	fmt.Printf("Brand: %s\n", c.brand)
	fmt.Printf("Model: %s\n", c.model)
	fmt.Printf("Year: %d\n", c.year)
	fmt.Printf("Price: $%.2f\n", c.price)
}
// func (c Car) setCarPrice(){  // updating structs need pointers otherwise it will copy the value and update the copied value.
// 	c.price = 20000.00
// }
func (c* car) setCarPrice(newPrice float64) string{ // NOTE: The concept of pointers is saperatly discussed in another tutorial.
	c.price = newPrice
	return "Price updated successfully!"
}


func main() {

	car1 := car{
		brand: "Toyota",
		model: "Camry",
		year: 2022,
		price: 25000.00,
	}
	fmt.Println(car1.setCarPrice(20000.00))
	car1.displayDetails()

	// Anonymous structs (Used only when you just want that one struct in the code and dont wanna create another insatance from it.)
	food := struct{
		name string
		price float64
		quantity int
		isVegetarian bool
	}{
		name: "Pizza",
		price: 10.99,
		quantity: 2,
		isVegetarian: false,
	}
	// avoid using anonymous structs.
	fmt.Println(food.name,food.price,food.quantity,food.isVegetarian)

	// Embedding structs
	truck1 := truck{
		car: car{
			brand: "Ford",
			model: "F-150",
			year: 2021,
			price: 35000.00,
		},
		capacity: 5000.00,
	}

	fmt.Println(truck1.brand) // dont have to do like "truck1.car.brand"
	fmt.Println(truck1.capacity)
	truck1.displayDetails()
}