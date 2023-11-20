# Aufgabe 3: Zauberschule

## Lösungsidee

Der Plan der Zauberschule kann als gewichteter Graph dargestellt werden, wobei die Knoten
je einem Feld des Plans entsprechen und die Kanten die möglichen Bewegungen zwischen den Feldern.
Das Gewicht einer Kante ist für alle Kanten, außer einem Stockwerkwechsel, 1. Für einen Stockwerkwechsel
beträgt es 3.
Um den schnellsten Weg zu finden, muss der Weg mit den geringsten Kosten gefunden werden,
wozu der ~~Djikstra-Algorithmus~~ Dijkstra-Algorithmus verwendet werden kann.
Für jeden Knoten wird die Distanz zum Startknoten gespeichert, die anfangs unendlich ist.
Die Distanz des Startknotens wird auf 0 gesetzt.
Dann wird der Knoten mit der geringsten Distanz, welcher mit der aktuellen Distanz noch nicht
besucht wurde, ausgewählt und die Distanz zu allen Nachbarn des Knotens wird aktualisiert,
falls die Distanz über den aktuellen Knoten kürzer ist. Dieser Prozess wird
so lange wiederholt, bis der Zielknoten erreicht wird. Dann ist der Weg mit den geringsten Kosten
gefunden.

## Umsetzung

Die Datei, welche als Argument übergeben wurde, wird in eine Datenstruktur umgewandelt,
welche den Plan als 3d-Array darstellt. Die erste Dimension ist die Stockwerknummer,
die zweite und dritte Dimension sind die x- und y-Koordinate des Feldes.
Dann 3d-Array enthält booleans die angeben, ob das Feld begehbar ist oder nicht.
Während des Einlesens wird auch die Start- und Zielposition gespeichert.
Der Plan wird nicht als Graph dargestellt, da die Kanten nicht gespeichert werden müssen.
Sie sind immer die gleiche, links, rechts, oben, unten und Stockwerkwechsel.
Es muss nur überprüft werden, ob das Feld begehbar ist.
Um die Distanz zu jedem Feld zu speichern, wird eine map `distances` erstellt.
Um später den Weg zu rekonstruieren, wird ein map `previous` erstellt, welche für jeden Knoten
den Vorgängerknoten speichert. Um den Knoten mit der geringsten Distanz zu finden, wird ein
Min-Queue `queue` erstellt, welcher nur die Knoten enthält, die noch nicht besucht mit ihrer
aktuellen Distanz besucht wurden.
Dann wird eine Schleife gestartet, welche so lange wiederholt wird,
bis der Zielknoten erreicht wurde oder die Queue leer ist.
In der Schleife wird der Knoten mit der geringsten Distanz aus der Queue entfernt.
Dann wird überprüft, ob der Knoten der Zielknoten ist, wenn ja, wird die Schleife beendet.
Dann wird für jeden Nachbarn die Distanz über den aktuellen Knoten berechnet.
Wenn die Distanz kürzer ist, als die aktuelle Distanz des Nachbarn,
oder es noch keine Distanz gibt, wird die Distanz des Nachbarn
auf die neue Distanz gesetzt und der Nachbar mit der neuen Distanz in die Queue eingefügt.
Der aktuelle Knoten wird als Vorgänger des Nachbarn gespeichert.
Nach dem Ende der Schleife wird der Weg rekonstruiert, indem der Vorgänger des Zielknotens
zum Weg hinzugefügt wird und dann sein Vorgänger usw., bis der Startknoten erreicht wurde.
Der Weg ist dann in umgekehrter Reihenfolge gespeichert, daher wird er umgedreht.

## Beispiele
```shell
$ zauberschule --no-route Beispieleingaben/Zauberschule/zauberschule0.txt
Weg als Liste, Start nicht inkludiert:
{Floor:1 X:5 Y:9}
{Floor:1 X:6 Y:9}
{Floor:1 X:7 Y:9}
{Floor:0 X:7 Y:9}

Der kürzeste Weg braucht 8 Sekunden.
```

```shell
$ zauberschule --no-route Beispieleingaben/Zauberschule/zauberschule1.txt
Weg als Liste, Start nicht inkludiert:
{Floor:0 X:12 Y:7}
{Floor:0 X:11 Y:7}
{Floor:0 X:11 Y:6}
{Floor:0 X:11 Y:5}

Der kürzeste Weg braucht 4 Sekunden.
```

```shell
$ zauberschule --no-route Beispieleingaben/Zauberschule/zauberschule2.txt
Weg als Liste, Start nicht inkludiert:
{Floor:0 X:22 Y:3}
{Floor:0 X:23 Y:3}
{Floor:1 X:23 Y:3}
{Floor:1 X:24 Y:3}
{Floor:1 X:25 Y:3}
{Floor:0 X:25 Y:3}
{Floor:0 X:25 Y:4}
{Floor:0 X:25 Y:5}
{Floor:0 X:26 Y:5}
{Floor:0 X:27 Y:5}

Der kürzeste Weg braucht 14 Sekunden.
```

```shell
$ zauberschule --no-route Beispieleingaben/Zauberschule/zauberschule3.txt
Weg als Liste, Start nicht inkludiert:
{Floor:0 X:3 Y:28}
{Floor:0 X:3 Y:29}
{Floor:0 X:4 Y:29}
{Floor:0 X:5 Y:29}
{Floor:0 X:5 Y:28}
{Floor:0 X:5 Y:27}
{Floor:0 X:6 Y:27}
{Floor:0 X:7 Y:27}
{Floor:0 X:8 Y:27}
{Floor:0 X:9 Y:27}
{Floor:0 X:9 Y:28}
{Floor:0 X:9 Y:29}
{Floor:0 X:10 Y:29}
{Floor:0 X:11 Y:29}
{Floor:0 X:12 Y:29}
{Floor:0 X:13 Y:29}
{Floor:0 X:13 Y:28}
{Floor:0 X:13 Y:27}
{Floor:0 X:14 Y:27}
{Floor:0 X:15 Y:27}
{Floor:0 X:16 Y:27}
{Floor:0 X:17 Y:27}
{Floor:0 X:18 Y:27}
{Floor:0 X:19 Y:27}
{Floor:0 X:19 Y:26}
{Floor:0 X:19 Y:25}
{Floor:0 X:20 Y:25}
{Floor:0 X:21 Y:25}

Der kürzeste Weg braucht 28 Sekunden.
```

```shell
$ zauberschule --no-route Beispieleingaben/Zauberschule/zauberschule4.txt
Weg als Liste, Start nicht inkludiert:
{Floor:0 X:73 Y:66}
{Floor:0 X:73 Y:65}
{Floor:0 X:72 Y:65}
{Floor:0 X:71 Y:65}
{Floor:1 X:71 Y:65}
{Floor:1 X:71 Y:64}
{Floor:1 X:71 Y:63}
{Floor:1 X:70 Y:63}
{Floor:1 X:69 Y:63}
{Floor:1 X:68 Y:63}
{Floor:1 X:67 Y:63}
{Floor:1 X:67 Y:62}
{Floor:1 X:67 Y:61}
{Floor:0 X:67 Y:61}
{Floor:0 X:66 Y:61}
{Floor:0 X:65 Y:61}
{Floor:0 X:64 Y:61}
{Floor:0 X:63 Y:61}
{Floor:0 X:62 Y:61}
{Floor:0 X:61 Y:61}
{Floor:0 X:60 Y:61}
{Floor:0 X:59 Y:61}
{Floor:0 X:59 Y:62}
{Floor:0 X:59 Y:63}
{Floor:0 X:58 Y:63}
{Floor:0 X:57 Y:63}
{Floor:0 X:57 Y:62}
{Floor:0 X:57 Y:61}
{Floor:0 X:56 Y:61}
{Floor:0 X:55 Y:61}
{Floor:0 X:54 Y:61}
{Floor:0 X:53 Y:61}
{Floor:0 X:52 Y:61}
{Floor:0 X:51 Y:61}
{Floor:0 X:50 Y:61}
{Floor:0 X:49 Y:61}
{Floor:1 X:49 Y:61}
{Floor:1 X:49 Y:62}
{Floor:1 X:49 Y:63}
{Floor:1 X:48 Y:63}
{Floor:1 X:47 Y:63}
{Floor:1 X:47 Y:62}
{Floor:1 X:47 Y:61}
{Floor:1 X:46 Y:61}
{Floor:1 X:45 Y:61}
{Floor:0 X:45 Y:61}
{Floor:0 X:45 Y:60}
{Floor:0 X:45 Y:59}
{Floor:0 X:44 Y:59}
{Floor:0 X:43 Y:59}
{Floor:0 X:43 Y:58}
{Floor:0 X:43 Y:57}
{Floor:0 X:42 Y:57}
{Floor:0 X:41 Y:57}
{Floor:0 X:40 Y:57}
{Floor:0 X:39 Y:57}
{Floor:0 X:39 Y:56}
{Floor:0 X:39 Y:55}
{Floor:0 X:38 Y:55}
{Floor:0 X:37 Y:55}
{Floor:1 X:37 Y:55}
{Floor:1 X:37 Y:56}
{Floor:1 X:37 Y:57}
{Floor:1 X:36 Y:57}
{Floor:1 X:35 Y:57}
{Floor:1 X:34 Y:57}
{Floor:1 X:33 Y:57}
{Floor:1 X:32 Y:57}
{Floor:1 X:31 Y:57}
{Floor:1 X:31 Y:56}
{Floor:1 X:31 Y:55}
{Floor:0 X:31 Y:55}

Der kürzeste Weg braucht 84 Sekunden.
```

```shell
$ zauberschule --no-route Beispieleingaben/Zauberschule/zauberschule5.txt
Weg als Liste, Start nicht inkludiert:
{Floor:1 X:13 Y:167}
{Floor:1 X:13 Y:168}
{Floor:1 X:13 Y:169}
{Floor:1 X:14 Y:169}
{Floor:1 X:15 Y:169}
{Floor:1 X:15 Y:170}
{Floor:1 X:15 Y:171}
{Floor:1 X:16 Y:171}
{Floor:1 X:17 Y:171}
{Floor:1 X:18 Y:171}
{Floor:1 X:19 Y:171}
{Floor:1 X:20 Y:171}
{Floor:1 X:21 Y:171}
{Floor:0 X:21 Y:171}
{Floor:0 X:22 Y:171}
{Floor:0 X:23 Y:171}
{Floor:0 X:24 Y:171}
{Floor:0 X:25 Y:171}
{Floor:0 X:26 Y:171}
{Floor:0 X:27 Y:171}
{Floor:0 X:28 Y:171}
{Floor:0 X:29 Y:171}
{Floor:0 X:29 Y:172}
{Floor:0 X:29 Y:173}
{Floor:1 X:29 Y:173}
{Floor:1 X:30 Y:173}
{Floor:1 X:31 Y:173}
{Floor:1 X:32 Y:173}
{Floor:1 X:33 Y:173}
{Floor:1 X:34 Y:173}
{Floor:1 X:35 Y:173}
{Floor:1 X:36 Y:173}
{Floor:1 X:37 Y:173}
{Floor:1 X:37 Y:172}
{Floor:1 X:37 Y:171}
{Floor:1 X:37 Y:170}
{Floor:1 X:37 Y:169}
{Floor:1 X:38 Y:169}
{Floor:1 X:39 Y:169}
{Floor:1 X:40 Y:169}
{Floor:1 X:41 Y:169}
{Floor:1 X:42 Y:169}
{Floor:1 X:43 Y:169}
{Floor:0 X:43 Y:169}
{Floor:0 X:43 Y:168}
{Floor:0 X:43 Y:167}
{Floor:0 X:43 Y:166}
{Floor:0 X:43 Y:165}
{Floor:0 X:43 Y:164}
{Floor:0 X:43 Y:163}
{Floor:0 X:43 Y:162}
{Floor:0 X:43 Y:161}
{Floor:1 X:43 Y:161}
{Floor:1 X:44 Y:161}
{Floor:1 X:45 Y:161}
{Floor:1 X:46 Y:161}
{Floor:1 X:47 Y:161}
{Floor:0 X:47 Y:161}
{Floor:0 X:48 Y:161}
{Floor:0 X:49 Y:161}
{Floor:0 X:50 Y:161}
{Floor:0 X:51 Y:161}
{Floor:0 X:51 Y:160}
{Floor:0 X:51 Y:159}
{Floor:0 X:51 Y:158}
{Floor:0 X:51 Y:157}
{Floor:0 X:52 Y:157}
{Floor:0 X:53 Y:157}
{Floor:0 X:53 Y:156}
{Floor:0 X:53 Y:155}
{Floor:0 X:53 Y:154}
{Floor:0 X:53 Y:153}
{Floor:0 X:54 Y:153}
{Floor:0 X:55 Y:153}
{Floor:1 X:55 Y:153}
{Floor:1 X:56 Y:153}
{Floor:1 X:57 Y:153}
{Floor:1 X:58 Y:153}
{Floor:1 X:59 Y:153}
{Floor:0 X:59 Y:153}
{Floor:0 X:59 Y:154}
{Floor:0 X:59 Y:155}
{Floor:0 X:59 Y:156}
{Floor:0 X:59 Y:157}
{Floor:0 X:60 Y:157}
{Floor:0 X:61 Y:157}
{Floor:0 X:62 Y:157}
{Floor:0 X:63 Y:157}
{Floor:0 X:64 Y:157}
{Floor:0 X:65 Y:157}
{Floor:1 X:65 Y:157}
{Floor:1 X:66 Y:157}
{Floor:1 X:67 Y:157}
{Floor:1 X:68 Y:157}
{Floor:1 X:69 Y:157}
{Floor:0 X:69 Y:157}
{Floor:0 X:70 Y:157}
{Floor:0 X:71 Y:157}
{Floor:0 X:72 Y:157}
{Floor:0 X:73 Y:157}
{Floor:0 X:73 Y:156}
{Floor:0 X:73 Y:155}
{Floor:0 X:72 Y:155}
{Floor:0 X:71 Y:155}

Der kürzeste Weg braucht 124 Sekunden.
```

## Quellcode
path/path.go
```go
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
```

plan/plan.go
```go
package plan

import (
	"BWINF/pkg/slice"
	"BWINF/zauberschule/coordinate"
)

type Plan struct {
	Start, End coordinate.Coordinate
	Walkable   [2][][]bool
}

func New(n, m int) *Plan {
	return &Plan{
		Walkable: [2][][]bool{
			slice.New2D[bool](n, m),
			slice.New2D[bool](n, m),
		},
	}
}

func (f *Plan) IsValid(c coordinate.Coordinate) bool {
	return c.Floor >= 0 && c.Floor < len(f.Walkable) &&
		c.Y >= 0 && c.Y < len(f.Walkable[c.Floor]) &&
		c.X >= 0 && c.X < len(f.Walkable[c.Floor][c.Y])
}

// IsWalkable returns true if the coordinate is valid and walkable
// (i.e., not a wall and not outside the map)
func (f *Plan) IsWalkable(c coordinate.Coordinate) bool {
	return f.IsValid(c) && f.Walkable[c.Floor][c.Y][c.X]
}

func (f *Plan) Set(c coordinate.Coordinate, b bool) {
	f.Walkable[c.Floor][c.Y][c.X] = b
}

func (f *Plan) Neighbors(c coordinate.Coordinate) []Neighbor {
	var neighbors []Neighbor

	for _, neighbor := range Neighbors(c) {
		if f.IsWalkable(neighbor.Coordinate) {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}
```

plan/neighbor.go
```go
package plan

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
```