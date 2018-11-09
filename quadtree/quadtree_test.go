package quadtree

import (
	"reflect"
	"testing"
)

func Test_aabbForSW(t *testing.T) {
	type args struct {
		boundary aabb
	}
	tests := []struct {
		name string
		args args
		want aabb
	}{
		{"match base", args{NewAABB(0, 0, 50, 50)}, NewAABB(-12.5, -12.5, 25, 25)},
		{"match base", args{NewAABB(-12.5, -12.5, 25, 25)}, NewAABB(-18.75, -18.75, 12.5, 12.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aabbForSW(tt.args.boundary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("aabbForSW() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_aabbForSE(t *testing.T) {
	type args struct {
		boundary aabb
	}
	tests := []struct {
		name string
		args args
		want aabb
	}{
		{"match base", args{NewAABB(0, 0, 50, 50)}, NewAABB(12.5, -12.5, 25, 25)},
		{"match base", args{NewAABB(12.5, -12.5, 25, 25)}, NewAABB(18.75, -18.75, 12.5, 12.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aabbForSE(tt.args.boundary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("aabbForSE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_aabbForNW(t *testing.T) {
	type args struct {
		boundary aabb
	}
	tests := []struct {
		name string
		args args
		want aabb
	}{
		{"match base", args{NewAABB(0, 0, 50, 50)}, NewAABB(-12.5, 12.5, 25, 25)},
		{"match base", args{NewAABB(-12.5, 12.5, 25, 25)}, NewAABB(-18.75, 18.75, 12.5, 12.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aabbForNW(tt.args.boundary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("aabbForNW() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_aabbForNE(t *testing.T) {
	type args struct {
		boundary aabb
	}
	tests := []struct {
		name string
		args args
		want aabb
	}{
		{"match base", args{NewAABB(0, 0, 50, 50)}, NewAABB(12.5, 12.5, 25, 25)},
		{"match base", args{NewAABB(12.5, 12.5, 25, 25)}, NewAABB(18.75, 18.75, 12.5, 12.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := aabbForNE(tt.args.boundary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("aabbForNE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuadtree_Count(t *testing.T) {
	type fields struct {
		boundary   aabb
		points     []XY
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
