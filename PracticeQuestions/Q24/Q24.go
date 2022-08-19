package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(6, 12)
	fmt.Printf("Distance between %+v and %+v = %0.2f", p1, p2, Distance(p1, p2))
}

type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func Distance(p1, p2 Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	distance := math.Sqrt(dx*dx + dy*dy)
	return distance
}
