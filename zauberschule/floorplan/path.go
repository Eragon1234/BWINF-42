package floorplan

import (
	"BWINF/pkg/heap"
	"BWINF/zauberschule/coordinate"
	"slices"
)

type Path []coordinate.Coordinate

func FindPath(plan *Floorplan) *Path {
	distances := make(map[coordinate.Coordinate]int)
	previous := make(map[coordinate.Coordinate]coordinate.Coordinate)

	queue := heap.New(func(a, b coordinate.Coordinate) bool {
		return distances[a] < distances[b]
	})

	distances[plan.Start] = 0
	queue.Push(plan.Start)

	for queue.Len() > 0 {
		currentNode := queue.Pop()
		if currentNode == plan.End {
			break
		}

		for _, neighbor := range plan.Neighbors(currentNode) {
			newDistance := distances[currentNode] + neighbor.Distance
			if currentDistance, ok := distances[neighbor.Coordinate]; !ok || newDistance < currentDistance {
				distances[neighbor.Coordinate] = newDistance
				previous[neighbor.Coordinate] = currentNode
				queue.Push(neighbor.Coordinate)
			}
		}
	}

	return buildPath(plan.Start, plan.End, previous)
}

func buildPath(start, end coordinate.Coordinate, previous map[coordinate.Coordinate]coordinate.Coordinate) *Path {
	var path Path

	for current := end; current != start; current = previous[current] {
		path = append(path, current)
	}

	slices.Reverse(path)

	return &path
}
