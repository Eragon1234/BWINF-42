package plan

import (
	"BWINF/pkg/slice"
	"BWINF/zauberschule/coordinate"
)

type Plan struct {
	Start, End coordinate.Coordinate
	Plan       [2][][]bool
}

func New(n, m int) *Plan {
	return &Plan{
		Plan: [2][][]bool{
			slice.New2D[bool](n, m),
			slice.New2D[bool](n, m),
		},
	}
}

func (f *Plan) IsValid(c coordinate.Coordinate) bool {
	return c.Floor >= 0 && c.Floor < len(f.Plan) &&
		c.Y >= 0 && c.Y < len(f.Plan[c.Floor]) &&
		c.X >= 0 && c.X < len(f.Plan[c.Floor][c.Y])
}

func (f *Plan) Get(c coordinate.Coordinate) bool {
	if !f.IsValid(c) {
		return false
	}
	return f.Plan[c.Floor][c.Y][c.X]
}

func (f *Plan) Set(c coordinate.Coordinate, b bool) {
	f.Plan[c.Floor][c.Y][c.X] = b
}

func (f *Plan) Neighbors(c coordinate.Coordinate) []Neighbor {
	var neighbors []Neighbor

	for _, neighbor := range Neighbors(c) {
		if f.Get(neighbor.Coordinate) {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}
