package floorplan

import (
	"BWINF/zauberschule/coordinate"
	"bufio"
	"fmt"
	"io"
)

func Parse(reader io.Reader) (*Floorplan, error) {
	scanner := bufio.NewScanner(reader)

	scanner.Scan()

	var n, m int
	_, err := fmt.Sscanf(scanner.Text(), "%d %d", &n, &m)
	if err != nil {
		return nil, err
	}

	plan := New(n, m)

	parseFloor(scanner, n, plan, 0)

	// skip the empty line
	scanner.Scan()

	parseFloor(scanner, n, plan, 1)

	return plan, nil
}

func parseFloor(scanner *bufio.Scanner, n int, plan *Floorplan, floor int) {
	for y := 0; y < n; y++ {
		scanner.Scan()
		for x, c := range scanner.Text() {
			pos := coordinate.New(floor, x, y)

			if c == 'A' {
				plan.Start = pos
			} else if c == 'B' {
				plan.End = pos
			}
			plan.Set(pos, c != '#')
		}
	}
}
