package quadtree

import (
	"reflect"
	"testing"
)

func Test_BoundaryForSW(t *testing.T) {
	type args struct {
		boundary Boundary
	}
	tests := []struct {
		name string
		args args
		want Boundary
	}{
		{"match base", args{NewBoundary(0, 0, 50, 50)}, NewBoundary(-12.5, -12.5, 25, 25)},
		{"match base", args{NewBoundary(-12.5, -12.5, 25, 25)}, NewBoundary(-18.75, -18.75, 12.5, 12.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoundaryForSW(tt.args.boundary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoundaryForSW() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BoundaryForSE(t *testing.T) {
	type args struct {
		boundary Boundary
	}
	tests := []struct {
		name string
		args args
		want Boundary
	}{
		{"match base", args{NewBoundary(0, 0, 50, 50)}, NewBoundary(12.5, -12.5, 25, 25)},
		{"match base", args{NewBoundary(12.5, -12.5, 25, 25)}, NewBoundary(18.75, -18.75, 12.5, 12.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoundaryForSE(tt.args.boundary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoundaryForSE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BoundaryForNW(t *testing.T) {
	type args struct {
		boundary Boundary
	}
	tests := []struct {
		name string
		args args
		want Boundary
	}{
		{"match base", args{NewBoundary(0, 0, 50, 50)}, NewBoundary(-12.5, 12.5, 25, 25)},
		{"match base", args{NewBoundary(-12.5, 12.5, 25, 25)}, NewBoundary(-18.75, 18.75, 12.5, 12.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoundaryForNW(tt.args.boundary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoundaryForNW() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BoundaryForNE(t *testing.T) {
	type args struct {
		boundary Boundary
	}
	tests := []struct {
		name string
		args args
		want Boundary
	}{
		{"match base", args{NewBoundary(0, 0, 50, 50)}, NewBoundary(12.5, 12.5, 25, 25)},
		{"match base", args{NewBoundary(12.5, 12.5, 25, 25)}, NewBoundary(18.75, 18.75, 12.5, 12.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoundaryForNE(tt.args.boundary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoundaryForNE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuadtree_Count(t *testing.T) {
	type fields struct {
		boundary   Boundary
		points     []Point
		northwest  *Quadtree
		northeast  *Quadtree
		southwest  *Quadtree
		southeast  *Quadtree
		subdivided bool
		level      int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quadtree{
				boundary:   tt.fields.boundary,
				points:     tt.fields.points,
				northwest:  tt.fields.northwest,
				northeast:  tt.fields.northeast,
				southwest:  tt.fields.southwest,
				southeast:  tt.fields.southeast,
				subdivided: tt.fields.subdivided,
				level:      tt.fields.level,
			}
			if got := q.Count(); got != tt.want {
				t.Errorf("Quadtree.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
