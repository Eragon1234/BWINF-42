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

func (c Coordinate) Add(other Coordinate) Coordinate {
	return Coordinate{
		Floor: c.Floor + other.Floor,
		X:     c.X + other.X,
		Y:     c.Y + other.Y,
	}
}
