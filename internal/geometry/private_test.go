package geometry

import (
	"math"
	"testing"
)

type expected_private struct {
	ok   bool
	x, y float64
}

func TestIntersection_private(t *testing.T) {
	expected := expected_private{true, 4, 4} // TODO

	edges := make_sun_rays_test_case()
	for i, edgeA := range edges {
		for j, edgeB := range edges {
			if i == j {
				continue
			}

			ok, x, y := intersection(simplify(edgeA, edgeB))
			result := expected_private{ok, x, y}
			if result != expected {
				t.Errorf("Invalid: %s AND %s: %v", edgeA, edgeB, result)
			}
		}
	}
}

func TestCross_private(t *testing.T) {
	expected := expected_private{true, 4, 4} // TODO

	edges := make_sun_rays_test_case()
	for i, edgeA := range edges {
		for j, edgeB := range edges {
			if i == j {
				continue
			}

			ok, x, y := cross(simplify(edgeA, edgeB))
			result := expected_private{ok, x, y}
			if result != expected {
				t.Errorf("Invalid: %s AND %s: %v", edgeA, edgeB, result)
			}
		}
	}
}

func TestInside_private(t *testing.T) {
	t.Error("TODO")
}

func BenchmarkCross_private(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = cross(0, 0, 5, 5, 0, 3, 5, 2)
	}
}

func BenchmarkIntersection_private(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _ = intersection(0, 0, 5, 5, 0, 3, 5, 2)
	}
}

func make_sun_rays_test_case() []*Line {
	const (
		angle  int     = 45
		radius float64 = 4
	)

	result := make([]*Line, 0, 360/angle)
	for i := 0; i < 360; i += angle {
		j := 180 + i
		if j > 360 {
			j -= 360
		}

		forward := float64(i) * (math.Pi / 180.0)
		revers := float64(j) * (math.Pi / 180.0)
		result = append(result, NewLine(
			NewPoint(
				radius+radius*math.Cos(revers),
				radius+radius*math.Sin(revers),
			),
			NewPoint(
				radius+radius*math.Cos(forward),
				radius+radius*math.Sin(forward),
			),
		))
	}

	return result
}
