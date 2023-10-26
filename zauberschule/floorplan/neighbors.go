package floorplan

import "BWINF/zauberschule/coordinate"

type Neighbor struct {
	coordinate.Coordinate
	Distance int
}

func NewNeighbor(coordinate coordinate.Coordinate, distance int) Neighbor {
	return Neighbor{
		Coordinate: coordinate,
		Distance:   distance,
	}
}

var (
	floorUp   = coordinate.New(1, 0, 0)
	floorDown = coordinate.New(-1, 0, 0)
	left      = coordinate.New(0, -1, 0)
	right     = coordinate.New(0, 1, 0)
	up        = coordinate.New(0, 0, -1)
	down      = coordinate.New(0, 0, 1)
)

func Neighbors(c coordinate.Coordinate) []Neighbor {
	return []Neighbor{
		NewNeighbor(c.Add(floorUp), 3),
		NewNeighbor(c.Add(floorDown), 3),
		NewNeighbor(c.Add(left), 1),
		NewNeighbor(c.Add(right), 1),
		NewNeighbor(c.Add(up), 1),
		NewNeighbor(c.Add(down), 1),
	}
}
