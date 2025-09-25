package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками на плоскости.
Точки представлены в виде структуры Point с инкапсулированными (приватными) полями x, y (типа float64) и конструктором.
Расстояние рассчитывается по формуле между координатами двух точек.

Подсказка: используйте функцию-конструктор NewPoint(x, y), Point и метод Distance(other Point) float64.
*/

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func main() {
	pointA := NewPoint(123.456, 6789)
	pointB := NewPoint(223.456, 3789)
	distance := pointA.Distance(pointB)
	fmt.Printf("Расстояние между точками A(%.2f, %.2f) и B(%.2f, %.2f) равно %.2f\n",
		pointA.x, pointA.y, pointB.x, pointB.y, distance)
}
