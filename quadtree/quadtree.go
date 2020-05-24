package quadtree

import "fmt"

const CAPACITY = 4

type Quadtree struct {
	boundary   Boundary
	points     []Point
	northwest  *Quadtree
	northeast  *Quadtree
	southwest  *Quadtree
	southeast  *Quadtree
	subdivided bool
	level      int
}

func NewQuadtree(boundary Boundary, level int) *Quadtree {
	quadtree := Quadtree{boundary: boundary,
		points:     make([]Point, 0, CAPACITY),
		subdivided: false,
		level:      level}

	return &quadtree
}

func (q *Quadtree) GetBoundary() Boundary {

	return q.boundary
}

func (q *Quadtree) Insert(point Point) bool {
	// check whether the point fits in the boundary
	if !q.boundary.ContainsPoint(point) {
		return false
	}
	// add the point if the capicity hasn't been reached
	if len(q.points) < CAPACITY {
		point.L = q.level
		q.points = append(q.points, point)
		return true
	}

	if !q.subdivided {
		q.subdivide()
		q.subdivided = true
	}
	// the point will only be inserted into one of these
	if q.northwest.Insert(point) {
		return true
	}
	if q.northeast.Insert(point) {
		return true
	}
	if q.southwest.Insert(point) {
		return true
	}
	if q.southeast.Insert(point) {
		return true
	}
	return true
}

// create four new boundaries
func (q *Quadtree) subdivide() {
	level := q.level + 1
	q.northeast = NewQuadtree(BoundaryForNE(q.boundary), level)
	q.northwest = NewQuadtree(BoundaryForNW(q.boundary), level)
	q.southeast = NewQuadtree(BoundaryForSE(q.boundary), level)
	q.southwest = NewQuadtree(BoundaryForSW(q.boundary), level)
}

func BoundaryForNE(boundary Boundary) Boundary {
	x := boundary.x + boundary.w/4
	y := boundary.y + boundary.h/4
	fmt.Println(x, y)
	return NewBoundary(x, y, boundary.w/2, boundary.h/2)
}

func BoundaryForNW(boundary Boundary) Boundary {
	x := boundary.x - boundary.w/4
	y := boundary.y + boundary.h/4

	return NewBoundary(x, y, boundary.w/2, boundary.h/2)
}

func BoundaryForSE(boundary Boundary) Boundary {
	x := boundary.x + boundary.w/4
	y := boundary.y - boundary.h/4

	return NewBoundary(x, y, boundary.w/2, boundary.h/2)
}

func BoundaryForSW(boundary Boundary) Boundary {
	x := boundary.x - boundary.w/4
	y := boundary.y - boundary.h/4

	return NewBoundary(x, y, boundary.w/2, boundary.h/2)
}

func (q *Quadtree) Query(search Boundary) []Point {

	pointsFound := make([]Point, 0)
	if !q.boundary.IntersectsBoundary(search) {
		return pointsFound
	}

	for i := 0; i < len(q.points); i++ {
		if search.ContainsPoint(q.points[i]) {
			pointsFound = append(pointsFound, q.points[i])
		}
	}

	if !q.subdivided {
		return pointsFound
	}
	pointsFound = append(pointsFound, q.northeast.Query(search)...)
	pointsFound = append(pointsFound, q.northwest.Query(search)...)
	pointsFound = append(pointsFound, q.southeast.Query(search)...)
	pointsFound = append(pointsFound, q.southwest.Query(search)...)

	return pointsFound
}

func (q *Quadtree) Count() int {

	count := 0

	if !q.subdivided {
		return 1
	}

	count = count + q.northeast.Count()
	count = count + q.northwest.Count()
	count = count + q.southeast.Count()
	count = count + q.southwest.Count()

	return count
}

func (q *Quadtree) All() []*Quadtree {
	quadtrees := make([]*Quadtree, 0)

	if !q.subdivided {
		quadtrees = append(quadtrees, q)
		return quadtrees
	}

	quadtrees = append(quadtrees, q.northeast.All()...)
	quadtrees = append(quadtrees, q.northwest.All()...)
	quadtrees = append(quadtrees, q.southeast.All()...)
	quadtrees = append(quadtrees, q.southwest.All()...)

	return quadtrees
}

func (q *Quadtree) Delete(point Point) {
	fmt.Println(q.boundary.ContainsPoint(point))
	if !q.boundary.ContainsPoint(point) {
		return
	}

	// does this Quadtree contain the point
	contained, _ := contains(q.points, point)
	if contained {
		points := q.Query(q.boundary)
		// Reset the properties of the Quadtrees
		q.subdivided = false
		q.northeast = nil
		q.northwest = nil
		q.southeast = nil
		q.southwest = nil
		q.points = make([]Point, 0)
		// reinsert the points
		for _, p := range points {
			if p != point {
				q.Insert(p)
			}
		}
		return
	}

	if q.subdivided {
		q.northeast.Delete(point)
		q.northwest.Delete(point)
		q.southeast.Delete(point)
		q.southwest.Delete(point)
	}

	return
}

func contains(s []Point, e Point) (bool, int) {
	for i, a := range s {
		if a == e {
			return true, i
		}
	}
	return false, -1
}
