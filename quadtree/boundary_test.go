package quadtree

import (
	"testing"
)

func Test_aabb_ContainsPoint(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		W float64
		H float64
	}
	type args struct {
		point Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"inside   ", fields{0, 0, 50, 50}, args{Point{0, 0, 0, 1}}, true},
		{"inside 1 ", fields{0, 0, 50, 50}, args{Point{10, 10, 0, 1}}, true},
		{"inside 2 ", fields{0, 0, 50, 50}, args{Point{-10, -10, 0, 1}}, true},
		{"inside 3 ", fields{0, 0, 50, 50}, args{Point{10, -10, 0, 1}}, true},
		{"inside 4 ", fields{0, 0, 50, 50}, args{Point{-10, 10, 0, 1}}, true},
		{"outside 1", fields{0, 0, 50, 50}, args{Point{26, 26, 0, 1}}, false},
		{"outside 2", fields{0, 0, 50, 50}, args{Point{-26, -26, 0, 1}}, false},
		{"outside 3", fields{0, 0, 50, 50}, args{Point{-26, 26, 0, 1}}, false},
		{"outside 4", fields{0, 0, 50, 50}, args{Point{26, -26, 0, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := Boundary{
				X: tt.fields.X,
				Y: tt.fields.Y,
				W: tt.fields.W,
				H: tt.fields.H,
			}
			if got := base.ContainsPoint(tt.args.point); got != tt.want {
				t.Errorf("boundary.ContainsPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Boundary_IntersectsBoundary(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		W float64
		H float64
	}
	type args struct {
		other Boundary
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"match base", fields{0, 0, 50, 50}, args{NewBoundary(0, 0, 50, 50)}, true},
		{"match base", fields{0, 0, 50, 50}, args{NewBoundary(-100, -100, 50, 50)}, false},
		{"match base", fields{0, 0, 50, 50}, args{NewBoundary(100, 100, 50, 50)}, false},
		{"match base", fields{0, 0, 50, 50}, args{NewBoundary(-100, -100, 200, 200)}, true},
		{"match base", fields{0, 0, 50, 50}, args{NewBoundary(100, 100, 200, 200)}, true},
		{"match base", fields{0, 0, 50, 50}, args{NewBoundary(-150, -150, 200, 200)}, false},
		{"match base", fields{0, 0, 50, 50}, args{NewBoundary(150, 150, 200, 200)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := Boundary{
				X: tt.fields.X,
				Y: tt.fields.Y,
				W: tt.fields.W,
				H: tt.fields.H,
			}
			if got := base.IntersectsBoundary(tt.args.other); got != tt.want {
				t.Errorf("Boundary.IntersectsBoundary() = %v, want %v", got, tt.want)
			}
		})
	}
}
