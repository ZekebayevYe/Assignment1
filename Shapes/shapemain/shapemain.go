package main

import (
	"fmt"

	"Assignment1/Shapes"
)

func main() {
	shapesList := []shapes.Shape{
		shapes.Rectangle{Length: 5, Width: 3},
		shapes.Circle{Radius: 4},
		shapes.Square{Length: 4},
		shapes.Triangle{SideA: 3, SideB: 4, SideC: 5},
	}

	for _, shape := range shapesList {
		shapes.PrintShapeDetails(shape)
		fmt.Println("_________________________________________________________--")
	}
}
