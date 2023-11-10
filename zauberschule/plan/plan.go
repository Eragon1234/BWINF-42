package plan

import (
	"BWINF/pkg/slice"
	"BWINF/zauberschule/coordinate"
)

type Plan struct {
	Start, End coordinate.Coordinate
	Walkable   [2][][]bool
}

func New(n, m int) *Plan {
	return &Plan{
		Walkable: [2][][]bool{
			slice.New2D[bool](n, m),
			slice.New2D[bool](n, m),
		},
	}
}

func (f *Plan) IsValid(c coordinate.Coordinate) bool {
	return c.Floor >= 0 && c.Floor < len(f.Walkable) &&
		c.Y >= 0 && c.Y < len(f.Walkable[c.Floor]) &&
		c.X >= 0 && c.X < len(f.Walkable[c.Floor][c.Y])
}

// IsWalkable returns true if the coordinate is valid and walkable
// (i.e., not a wall and not outside the map)
func (f *Plan) IsWalkable(c coordinate.Coordinate) bool {
	return f.IsValid(c) && f.Walkable[c.Floor][c.Y][c.X]
}

func (f *Plan) Set(c coordinate.Coordinate, b bool) {
	f.Walkable[c.Floor][c.Y][c.X] = b
}

func (f *Plan) Neighbors(c coordinate.Coordinate) []Neighbor {
	var neighbors []Neighbor

	for _, neighbor := range Neighbors(c) {
		if f.IsWalkable(neighbor.Coordinate) {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}
