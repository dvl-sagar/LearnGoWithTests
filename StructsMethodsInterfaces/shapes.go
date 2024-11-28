package StructMethodsInterface

type Rectangle struct{
	breadth float64
	length float64
}

func Perimeter(rect Rectangle) float64{
	return 2*(rect.breadth + rect.length)
}

func Area(rect Rectangle)float64{
	return rect.breadth * rect.length
}