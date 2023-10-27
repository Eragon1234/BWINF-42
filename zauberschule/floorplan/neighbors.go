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

func Neighbors(c coordinate.Coordinate) []Neighbor {
	return []Neighbor{
		NewNeighbor(c.Add(coordinate.FloorUp), 3),
		NewNeighbor(c.Add(coordinate.FloorDown), 3),
		NewNeighbor(c.Add(coordinate.Left), 1),
		NewNeighbor(c.Add(coordinate.Right), 1),
		NewNeighbor(c.Add(coordinate.Up), 1),
		NewNeighbor(c.Add(coordinate.Down), 1),
	}
}
