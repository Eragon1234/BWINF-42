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
