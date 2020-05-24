package quadtree

type Boundary struct {
	X float64 // minimum longitude
	Y float64 // minimum latiude
	W float64 // degrees east from minimum longitude
	H float64 // degress north from minimum latitude
}

// crate a new aabb rectangle
func NewBoundary(x float64, y float64, w float64, h float64) Boundary {
	boundary := Boundary{
		X: x,
		Y: y,
		W: w,
		H: h,
	}
	return boundary
}

func (base Boundary) ContainsPoint(point Point) bool {

	return ((point.X < base.X+base.W/2 && point.X >= base.X-base.W/2) &&
		(point.Y < base.Y+base.H/2 && point.Y >= base.Y-base.H/2))
}

func (base Boundary) IntersectsBoundary(other Boundary) bool {

	if (other.X+other.W/2 >= base.X-base.W/2 &&
		other.X-other.W/2 < base.X+base.W/2) &&
		(other.Y+other.H/2 >= base.Y-base.H/2 &&
			other.Y-other.H/2 < base.Y+base.H/2) {
		return true
	}
	return false
}
