package geometry

import "fmt"

type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func (point *Point) String() string {
	return fmt.Sprintf("{ %f, %f }", point.X, point.Y)
}
