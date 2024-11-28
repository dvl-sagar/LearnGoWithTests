package StructMethodsInterface

import "math"

type Rectangle struct {
	breadth float64
	length  float64
}

type Circle struct {
	Radius float64
}

type Shape interface{
	Area() float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.breadth + rect.length)
}

func (r Rectangle) Area() float64 {
	return r.breadth * r.length
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}