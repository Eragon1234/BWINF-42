package path

import (
	"BWINF/pkg/heap"
	"BWINF/zauberschule/coordinate"
	"BWINF/zauberschule/plan"
	"slices"
)

type Path []coordinate.Coordinate

func FindPath(p *plan.Plan) (*Path, int) {
	distances := make(map[coordinate.Coordinate]int)
	previous := make(map[coordinate.Coordinate]coordinate.Coordinate)

	queue := heap.New(func(a, b coordinate.Coordinate) bool {
		return distances[a] < distances[b]
	})

	distances[p.Start] = 0
	queue.Push(p.Start)

	for queue.Len() > 0 {
		currentNode := queue.Pop()
		if currentNode == p.End {
			break
		}

		for _, neighbor := range p.Neighbors(currentNode) {
			newDistance := distances[currentNode] + neighbor.Distance
			if currentDistance, ok := distances[neighbor.Coordinate]; !ok || newDistance < currentDistance {
				distances[neighbor.Coordinate] = newDistance
				previous[neighbor.Coordinate] = currentNode
				queue.Push(neighbor.Coordinate)
			}
		}
	}

	return buildPath(p.Start, p.End, previous), distances[p.End]
}

func buildPath(start, end coordinate.Coordinate, previous map[coordinate.Coordinate]coordinate.Coordinate) *Path {
	var path Path

	for current := end; current != start; current = previous[current] {
		path = append(path, current)
	}

	slices.Reverse(path)

	return &path
}
