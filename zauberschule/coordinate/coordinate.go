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

var (
	Up        = New(0, 0, -1)
	Down      = New(0, 0, 1)
	Left      = New(0, -1, 0)
	Right     = New(0, 1, 0)
	FloorUp   = New(-1, 0, 0)
	FloorDown = New(1, 0, 0)
)

func (c Coordinate) Add(other Coordinate) Coordinate {
	return New(c.Floor+other.Floor, c.X+other.X, c.Y+other.Y)
}

func (c Coordinate) Sub(other Coordinate) Coordinate {
	return New(c.Floor-other.Floor, c.X-other.X, c.Y-other.Y)
}
