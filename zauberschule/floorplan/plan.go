package floorplan

import (
	"BWINF/pkg/slice"
)

type Floorplan struct {
	Start, End [3]int
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
