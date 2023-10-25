package floorplan

import (
	"BWINF/pkg/slice"
	"BWINF/zauberschule/coordinate"
)

type Floorplan struct {
	Start, End coordinate.Coordinate
	Plan       [2][][]bool
}

func New(n, m int) *Floorplan {
	return &Floorplan{
		Plan: [2][][]bool{
			slice.New2D[bool](n, m),
			slice.New2D[bool](n, m),
		},
	}
}

func (f *Floorplan) IsValid(c coordinate.Coordinate) bool {
	return c.Floor >= 0 && c.Floor < len(f.Plan) &&
		c.Y >= 0 && c.Y < len(f.Plan[c.Floor]) &&
		c.X >= 0 && c.X < len(f.Plan[c.Floor][c.Y])
}

func (f *Floorplan) Get(c coordinate.Coordinate) bool {
	if !f.IsValid(c) {
		return false
	}
	return f.Plan[c.Floor][c.Y][c.X]
}

func (f *Floorplan) Set(c coordinate.Coordinate, b bool) {
	f.Plan[c.Floor][c.Y][c.X] = b
}
