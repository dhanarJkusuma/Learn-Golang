package car

import (
	"fmt"
)

type Car struct {
	Brand string
	CoordinateX int
	CoordinateY int
}

func GenerateNewCar(b string, x int, y int) *Car {
	return &Car{
		Brand: b,
		CoordinateX: x,
		CoordinateY: y,
	}
}

func MoveRight(car *Car, distance int){
	fmt.Print("Moving car to `Right direction` (%d)", distance)
	car.CoordinateX += distance
}

func MoveLeft(car *Car, distance int){
	fmt.Print("Moving car to `Left direction` (%d)", distance)
	car.CoordinateX -= distance
}

func MoveTop(car *Car, distance int){
	fmt.Print("Moving car to `Top direction` (%d)", distance)
	car.CoordinateY += distance
}

func MoveBottom(car *Car, distance int){
	fmt.Print("Moving car to `Bottom direction` (%d)", distance)
	car.CoordinateY -= distance
}

func ShowCoordinate(car Car){
	fmt.Print("X:Y => %d:%d", car.CoordinateX, car.CoordinateY)
}

