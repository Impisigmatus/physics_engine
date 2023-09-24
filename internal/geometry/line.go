package geometry

import (
	"fmt"
	"math"
)

type Line struct {
	Begin, End *Point
	Length     float64
}

func NewLine(begin, end *Point) *Line {
	deltaX := math.Abs(end.X - begin.X)
	deltaY := math.Abs(end.Y - begin.Y)

	return &Line{
		Begin:  begin,
		End:    end,
		Length: math.Sqrt(deltaX*deltaX + deltaY*deltaY),
	}
}

func (line *Line) String() string {
	return fmt.Sprintf("{ 'begin': %s, 'end': %s, 'length': %f }", line.Begin, line.End, line.Length)
}
