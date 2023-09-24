package geometry

import "fmt"

type Rectangle struct {
	LineA *Line
	LineB *Line
	LineC *Line
	LineD *Line
}

// begin - is top left point
// end -  is bottom rigth point
func NewRectangle(begin, end *Point) Geometry {
	lengthX := end.X - begin.X
	lengthY := end.Y - begin.Y

	return &Rectangle{
		LineA: NewLine(begin, NewPoint(begin.X+lengthX, begin.Y)),
		LineB: NewLine(NewPoint(begin.X+lengthX, begin.Y), end),
		LineC: NewLine(end, NewPoint(begin.X, begin.Y+lengthY)),
		LineD: NewLine(NewPoint(begin.X, begin.Y+lengthY), begin),
	}
}

func (rectangle *Rectangle) Intersect(geom Geometry) bool {
	return LinesIntersection(rectangle.GetEdges(), geom.GetEdges())
}

func (rectangle *Rectangle) GetBoundingBox() *Rectangle {
	return rectangle
}

func (rectangle *Rectangle) GetEdges() []*Line {
	if !rectangle.IsValid() {
		return nil
	}

	return []*Line{
		rectangle.LineA,
		rectangle.LineB,
		rectangle.LineC,
		rectangle.LineD,
	}
}

func (rectangle *Rectangle) IsValid() bool {
	if rectangle.LineA == nil || rectangle.LineB == nil || rectangle.LineC == nil || rectangle.LineD == nil {
		return false
	}

	return true
}

func (rectangle *Rectangle) String() string {
	return fmt.Sprintf("{ %s }", rectangle.GetEdges())
}
