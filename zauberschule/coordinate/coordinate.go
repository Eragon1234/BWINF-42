package coordinate

type Coordinate struct {
	Floor int
	X, Y  int
}

func New(floor, x, y int) Coordinate {
	return Coordinate{
		Floor: floor,
		X:     x,
		Y:     y,
	}
}
