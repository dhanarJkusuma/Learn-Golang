package main

import (
	"fmt"
	"hello-go/car" 
)



func main(){

	bmw := car.GenerateNewCar("BMW", 0, 0)
	fmt.Print("Starting engine new car : ", bmw.Brand)

	car.MoveTop(bmw, 10)
	car.MoveRight(bmw, 5)
	car.MoveBottom(bmw, 1)
	car.ShowCoordinate(*bmw)
}



