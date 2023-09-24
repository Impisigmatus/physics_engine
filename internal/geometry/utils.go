package geometry

func LineIntersection(edgeA, edgeB *Line) *Point {
	ok, x, y := intersection(simplify(edgeA, edgeB))
	if !ok {
		return nil
	}

	return NewPoint(x, y)
}

func LineCross(edgeA, edgeB *Line) *Point {
	ok, x, y := cross(simplify(edgeA, edgeB))
	if !ok {
		return nil
	}

	return NewPoint(x, y)
}

func LinesIntersection(a, b []*Line) bool {
	for _, edgeA := range a {
		for _, edgeB := range b {
			if point := LineIntersection(edgeA, edgeB); point != nil {
				return true
			}
		}
	}

	return false
}

func IsClosed(edges []Line) bool {
	for i, length := 0, len(edges); i < length; i++ {
		next := i + 1
		if next == length {
			next = 0
		}

		if edges[i].End != edges[next].Begin {
			return false
		}
	}

	return true
}
