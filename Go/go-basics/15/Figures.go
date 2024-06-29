package main

import (
	"fmt"
	"math"
)

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

var constsNames = map[int]string{
	0: "Square",
	1: "Circle",
	2: "Triangle",
}

func main() {
	var myFigure figures = 2

	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("Ошибка")
		return
	}
	x := 5.0

	myArea := ar(x)
	fmt.Printf("Площадь фигуры %v = %f", constsNames[int(myFigure)], myArea)
}
func area(f figures) (func(float64) float64, bool) {
	switch f {
	case square:
		return func(f float64) float64 { return f * f }, true
	case circle:
		return func(f float64) float64 { return math.Pi * math.Pow(f, 2) }, true
	case triangle:
		return func(f float64) float64 { return math.Sqrt(3) * math.Pow(f, 2) / 4 }, true
	default:
		return nil, false
	}
}
