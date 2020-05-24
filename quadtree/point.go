package quadtree

// Point is a Point in the Quadtree
type Point struct {
	X float64
	Y float64
	L int   // which Level of the quadtree is the point in
	I int64 // id
}
