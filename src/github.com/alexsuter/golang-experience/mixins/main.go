package main

import (
	"fmt"
)

type Car struct {
	wheelCount int
}

func (car Car) numberOfWheels() int {
	return car.wheelCount
}

type Porsche struct {
	Car
}

func (porsche Porsche) numberOfWheels() int {
	return 0
}

func main() {
	var p = Porsche{Car{4}}
	
	// Use overridden method
	fmt.Println(p.numberOfWheels())
	
	// Use super method
	fmt.Println(p.Car.numberOfWheels())
}
