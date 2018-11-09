package quadtree

type aabb struct {
	x float64
	y float64
	w float64
	h float64
}

type AABB_JSON struct {
	X float64
	Y float64
	W float64
	H float64
}

func NewAABB(x float64, y float64, w float64, h float64) aabb {
	aabb := aabb{
		x: x,
		y: y,
		w: w,
		h: h,
	}

	return aabb
}

func (base aabb) GetJSON() AABB_JSON {

	return AABB_JSON{base.x, base.y, base.w, base.h}
}

func (base aabb) ContainsPoint(point XY) bool {

	return ((point.X < base.x+base.w/2 && point.X >= base.x-base.w/2) &&
		(point.Y < base.y+base.h/2 && point.Y >= base.y-base.h/2))
}

func (base aabb) IntersectsAABB(other aabb) bool {

	if (other.x+other.w/2 >= base.x-base.w/2 &&
		other.x-other.w/2 < base.x+base.w/2) &&
		(other.y+other.h/2 >= base.y-base.h/2 &&
			other.y-other.h/2 < base.y+base.h/2) {
		return true
	}

	return false
}
