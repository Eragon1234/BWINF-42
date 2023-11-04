package route

import (
	"BWINF/pkg/ansi"
	"BWINF/zauberschule/coordinate"
	"BWINF/zauberschule/path"
	"BWINF/zauberschule/plan"
	"slices"
	"strings"
)

var moveToString = map[coordinate.Coordinate]string{
	coordinate.Up:        "↑",
	coordinate.Down:      "↓",
	coordinate.Left:      "←",
	coordinate.Right:     "→",
	coordinate.FloorUp:   "\u219F",
	coordinate.FloorDown: "\u21A1",
}

type printPath struct {
	Coordinate coordinate.Coordinate
	Symbol     string
}

func StringFloorplanWithPath(f plan.Plan, path path.Path) string {
	styles := make(map[coordinate.Coordinate]ansi.Style)

	p := constructStepsAndColors(path, f, styles)

	sb := new(strings.Builder)
	writeFloor(sb, f, &p, 0, styles)
	sb.WriteRune('\n')
	writeFloor(sb, f, &p, 1, styles)
	sb.WriteRune('\n')

	return sb.String()
}

func constructStepsAndColors(p path.Path, f plan.Plan, styles map[coordinate.Coordinate]ansi.Style) []printPath {
	possibleColors := []ansi.Style{
		ansi.Red,
		ansi.Green,
		ansi.Yellow,
		ansi.Blue,
		ansi.Magenta,
		ansi.Cyan,
		ansi.RGB(255, 165, 0), // orange
		ansi.RGB(149, 11, 22), // some red
		ansi.RGB(112, 196, 141),
		ansi.RGB(81, 148, 6), // some green
		ansi.RGB(228, 245, 234),
	}

	p = append(path.Path{f.Start}, p...)

	var printPaths []printPath

	styles[f.Start] = styles[f.Start].Add(ansi.Underline)

	for i, c := range p {
		if i == len(p)-1 {
			break
		}
		next := p[i+1]
		move := next.Sub(c)
		printPaths = append(printPaths, printPath{
			Coordinate: c,
			Symbol:     moveToString[move],
		})

		if move == coordinate.FloorUp || move == coordinate.FloorDown {
			nextColor := possibleColors[0]
			styles[c] = styles[c].Add(nextColor)
			styles[next] = styles[next].Add(nextColor)
			possibleColors = possibleColors[1:]
		}
	}

	printPaths = append(printPaths, printPath{
		Coordinate: f.End,
		Symbol:     "X",
	})

	slices.SortFunc(printPaths, func(a, b printPath) int {
		if a.Coordinate.Floor != b.Coordinate.Floor {
			return a.Coordinate.Floor - b.Coordinate.Floor
		}
		if a.Coordinate.Y != b.Coordinate.Y {
			return a.Coordinate.Y - b.Coordinate.Y
		}
		return a.Coordinate.X - b.Coordinate.X
	})

	return printPaths
}

var (
	emptyField = ansi.S(".", ansi.RGB(130, 130, 130))
	wall       = ansi.S("#", ansi.RGB(130, 130, 130))
)

func writeFloor(sb *strings.Builder, f plan.Plan, path *[]printPath, floor int, colors map[coordinate.Coordinate]ansi.Style) {
	for y, row := range f.Plan[floor] {
		for x, node := range row {
			c := coordinate.New(floor, x, y)

			var s string

			if step, ok := getStep(path, c); ok {
				s = step.Symbol
			} else if node {
				s = emptyField
			} else {
				s = wall
			}

			color, ok := colors[c]
			if ok {
				s = ansi.S(s, color)
			} else {
			}

			sb.WriteString(s)
		}
		sb.WriteRune('\n')
	}
}

// getStep returns the first element in path if there is one, and it's equal to c
func getStep(path *[]printPath, c coordinate.Coordinate) (printPath, bool) {
	if len(*path) == 0 {
		return printPath{}, false
	}
	p := (*path)[0]

	if p.Coordinate != c {
		return printPath{}, false
	}

	*path = (*path)[1:]

	return p, true
}
