package quadtree

type Boundary struct {
	x float64 // minimum longitude
	y float64 // minimum latiude
	w float64 // degrees east from minimum longitude
	h float64 // degress north from minimum latitude
}

type Boundary_json struct {
	X float64
	Y float64
	W float64
	H float64
}

// crate a new aabb rectangle
func NewBoundary(x float64, y float64, w float64, h float64) Boundary {
	boundary := Boundary{
		x: x,
		y: y,
		w: w,
		h: h,
	}
	return boundary
}

func (base Boundary) GetJSON() Boundary_json {

	return Boundary_json{base.x, base.y, base.w, base.h}
}

func (base Boundary) ContainsPoint(point Point) bool {

	return ((point.X < base.x+base.w/2 && point.X >= base.x-base.w/2) &&
		(point.Y < base.y+base.h/2 && point.Y >= base.y-base.h/2))
}

func (base Boundary) IntersectsBoundary(other Boundary) bool {

	if (other.x+other.w/2 >= base.x-base.w/2 &&
		other.x-other.w/2 < base.x+base.w/2) &&
		(other.y+other.h/2 >= base.y-base.h/2 &&
			other.y-other.h/2 < base.y+base.h/2) {
		return true
	}
	return false
}
