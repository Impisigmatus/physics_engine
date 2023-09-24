package geometry

type Geometry interface {
	Intersect(geom Geometry) bool
	GetBoundingBox() *Rectangle
	GetEdges() []*Line

	IsValid() bool
	String() string
}
