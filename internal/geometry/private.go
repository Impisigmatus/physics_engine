package geometry

import "math"

// Упрощение
func simplify(edgeA, edgeB *Line) (x1, y1, x2, y2, x3, y3, x4, y4 float64) {
	x1 = edgeA.Begin.X
	y1 = edgeA.Begin.Y
	x2 = edgeA.End.X
	y2 = edgeA.End.Y

	x3 = edgeB.Begin.X
	y3 = edgeB.Begin.Y
	x4 = edgeB.End.X
	y4 = edgeB.End.Y

	return x1, y1, x2, y2, x3, y3, x4, y4
}

// Пересечение отрезков
func intersection(x1, y1, x2, y2, x3, y3, x4, y4 float64) (ok bool, x float64, y float64) {
	ok, x, y = cross(x1, y1, x2, y2, x3, y3, x4, y4)
	if !ok {
		return false, 0, 0
	}

	// TODO: объединить две области в одну путем поиска области пересечения
	if inside(x, y, x1, y1, x2, y2) && inside(x, y, x3, y3, x4, y4) {
		return true, x, y
	}

	return false, 0, 0
}

// Точка внутри рамки
func inside(x, y, x1, y1, x2, y2 float64) bool {
	if x >= x1 && x <= x2 && y >= y1 && x <= y2 {
		return true
	}

	return false
}

// Пересечение лучей
func cross(x1, y1, x2, y2, x3, y3, x4, y4 float64) (ok bool, x float64, y float64) {
	if x1 == x4 && x2 == x3 || y1 == y4 && y2 == y3 {
		return true,
			round(x1 + (x2-x1)/2),
			round(y1 + (y2-y1)/2)
	}

	var n float64
	if y2-y1 != 0 { // a(y)
		q := (x2 - x1) / (y1 - y2)
		sn := (x3 - x4) + (y3-y4)*q // c(x) + c(y)*q
		if sn == 0 {
			return false, 0, 0
		}

		fn := (x3 - x1) + (y3-y1)*q // b(x) + b(y)*q
		n = fn / sn
	} else {
		if y4-y3 != 0 { // b(y)
			return false, 0, 0
		}

		n = (y3 - y1) / (y3 - y4) // c(y)/b(y)
	}

	return true,
		round(x3 + (x4-x3)*n), // x3 + (-b(x))*n
		round(y3 + (y4-y3)*n) // y3 + (-b(y))*n
}

func round(x float64) float64 {
	const precision = 100000000000000
	return math.Round(x*precision) / precision
}
