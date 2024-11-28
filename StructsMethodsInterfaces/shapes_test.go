package StructMethodsInterface

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	testCases := []struct{
		shape Shape
		want float64
	}{
		{Rectangle{10,10},100},
		{Circle{10},314.1592653589793},
		{Triangle{10, 6}, 30.0},
	}
	for _,tt := range testCases{
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
