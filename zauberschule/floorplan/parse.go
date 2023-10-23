package floorplan

import (
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
	for i := 0; i < n; i++ {
		scanner.Scan()
		for j, c := range scanner.Text() {
			if c == 'A' {
				plan.Start = [3]int{floor, i, j}
			} else if c == 'B' {
				plan.End = [3]int{floor, i, j}
			}
			plan.Plan[floor][i][j] = c != '#'
		}
	}
}
