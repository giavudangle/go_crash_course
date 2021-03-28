package main

import (
	"fmt"
	"math"
)

// Define interfaces

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{
		0,
		0,
		2,
	}

	rectangle := Rectangle{
		12,
		5,
	}

	fmt.Println(getArea(circle))
	fmt.Println(getArea(rectangle))

}
