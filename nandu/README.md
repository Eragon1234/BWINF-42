# Aufgabe 4: Nandu

## Lösungsidee

Eine Konstruktion aus den Blöcken kann als ein Graph dargestellt werden,
bei dem die Knoten die Blöcke sind und die Kanten eine Verbindung einer LED
zu einem Sensor darstellt. Die Kanten besitzen einen Wert der darstellt, ob die LED
an ist.
Zur Simulation einer möglichen Eingabe erhalten die Eingabe Kanten
ihren Eingabewert und basierend auf der Farbe des Blocks,
zu dem sie führen, wird der Wert der ausgehenden Kanten dieses Blockes
berechnet. Dies wird wiederholt bis das Ende des Graphen erreicht ist.
Die Ausgabe Kanten besitzen dann den Wert, wie die Konstruktion auf diese Eingaben reagieren würde.
Um die Ausgabe für sämtliche Eingaben zu berechnen, wird der Graph für jede mögliche Eingabe simuliert.

## Umsetzung

Die Blöcke werden als ein struct `Node` dargestellt,
sie besitzen einen Ein- und einen Ausgabewert dargestellt als `bool`,
einen Pointer auf den nächsten Block und eine Funktion.
Die Blöcke stellen nur einen der zwei Teile eines Blocks dar.
Um auf den Eingabewert ihres zugehörigen Blocks zuzugreifen,
besitzen sie einen Pointer auf den Block der ihren Partner darstellt.
Die Funktion eines Blocks erhält den Eingabewert des Blocks und den
Eingabewert des Partners und gibt den Ausgabewert zurück.
Für einen weiß gefärbten Block wird die Negation der Konjunktion der Eingabewerte
zurückgegeben. Für einen blauen Block wird der Eingabewert des eigenen Blocks zurückgegeben.
Für einen roten Block gibt es zwei Funktionen,
die Funktion für den Teil eines roten Blocks mit Sensor gibt die Negation des Eingabewertes
des eigenen Blocks zurück, die Funktion für den Teil eines roten Blocks ohne Sensor
gibt die Negation des Eingabewertes des Partners zurück.
Diese Blöcke besitzen eine Funktion `Set`, mit welcher ihr Eingabewert gesetzt werden kann.
Wenn der neue Eingabewert nicht dem alten entspricht, wird der neue Ausgabewert
berechnet und der Wert des Partners wird auch aktualisiert.
Dann wird die Funktion `Set` des nächsten Blocks mit dem Ausgabewert aufgerufen.
Falls es keinen nächsten Block gibt, wird nichts getan.
Dadurch werden alle Werte des Graphen rekursiv aktualisiert.
Die Eingabeblöcke und Ausgabeblöcke werden als Block dargestellt,
ohne einen Partner und mit einer Funktion, die den Eingabewert des Blocks zurückgibt.
Um die Ausgaben für alle möglichen Eingaben zu erhalten,
wird der Graph für jede mögliche Eingabe simuliert.
Um alle möglichen Eingaben zu erhalten wird von 0 bis 2<sup>Anzahl der Eingaben</sup> gezählt
und die Bits der Zahl werden als Eingabewerte verwendet.

## Beispiele

```shell
$ nandu Beispieleingaben/Nandu/nandu1.txt
Q1: false
Q2: false
L1: true
L2: true

Q1: true
Q2: false
L1: true
L2: true

Q1: false
Q2: true
L1: true
L2: true

Q1: true
Q2: true
L1: false
L2: false
```

```shell
$ nandu Beispieleingaben/Nandu/nandu2.txt
Q1: false
Q2: false
L1: false
L2: true

Q1: true
Q2: false
L1: false
L2: true

Q1: false
Q2: true
L1: false
L2: true

Q1: true
Q2: true
L1: true
L2: false
```

```shell
$ nandu Beispieleingaben/Nandu/nandu3.txt
Q1: false
Q2: false
Q3: false
L1: true
L2: false
L3: false
L4: true

Q1: true
Q2: false
Q3: false
L1: false
L2: true
L3: false
L4: true

Q1: false
Q2: true
Q3: false
L1: true
L2: false
L3: true
L4: true

Q1: true
Q2: true
Q3: false
L1: false
L2: true
L3: true
L4: true

Q1: false
Q2: false
Q3: true
L1: true
L2: false
L3: false
L4: false

Q1: true
Q2: false
Q3: true
L1: false
L2: true
L3: false
L4: false

Q1: false
Q2: true
Q3: true
L1: true
L2: false
L3: true
L4: false

Q1: true
Q2: true
Q3: true
L1: false
L2: true
L3: true
L4: false
```

```shell
$ nandu Beispieleingaben/Nandu/nandu4.txt
Q1: false
Q2: false
Q3: false
Q4: false
L1: false
L2: false

Q1: true
Q2: false
Q3: false
Q4: false
L1: false
L2: false

Q1: false
Q2: true
Q3: false
Q4: false
L1: true
L2: false

Q1: true
Q2: true
Q3: false
Q4: false
L1: false
L2: false

Q1: false
Q2: false
Q3: true
Q4: false
L1: false
L2: true

Q1: true
Q2: false
Q3: true
Q4: false
L1: false
L2: true

Q1: false
Q2: true
Q3: true
Q4: false
L1: true
L2: true

Q1: true
Q2: true
Q3: true
Q4: false
L1: false
L2: true

Q1: false
Q2: false
Q3: false
Q4: true
L1: false
L2: false

Q1: true
Q2: false
Q3: false
Q4: true
L1: false
L2: false

Q1: false
Q2: true
Q3: false
Q4: true
L1: true
L2: false

Q1: true
Q2: true
Q3: false
Q4: true
L1: false
L2: false

Q1: false
Q2: false
Q3: true
Q4: true
L1: false
L2: false

Q1: true
Q2: false
Q3: true
Q4: true
L1: false
L2: false

Q1: false
Q2: true
Q3: true
Q4: true
L1: true
L2: false

Q1: true
Q2: true
Q3: true
Q4: true
L1: false
L2: false
```

```shell
$ nandu Beispieleingaben/Nandu/nandu5.txt
Q1: false
Q2: false
Q3: false
Q4: false
Q5: false
Q6: false
L1: false
L2: false
L3: false
L4: true
L5: false

Q1: true
Q2: false
Q3: false
Q4: false
Q5: false
Q6: false
L1: true
L2: false
L3: false
L4: true
L5: false

Q1: false
Q2: true
Q3: false
Q4: false
Q5: false
Q6: false
L1: false
L2: false
L3: false
L4: true
L5: false

Q1: true
Q2: true
Q3: false
Q4: false
Q5: false
Q6: false
L1: true
L2: false
L3: false
L4: true
L5: false

Q1: false
Q2: false
Q3: true
Q4: false
Q5: false
Q6: false
L1: false
L2: false
L3: false
L4: true
L5: false

Q1: true
Q2: false
Q3: true
Q4: false
Q5: false
Q6: false
L1: true
L2: false
L3: false
L4: true
L5: false

Q1: false
Q2: true
Q3: true
Q4: false
Q5: false
Q6: false
L1: false
L2: false
L3: false
L4: true
L5: false

Q1: true
Q2: true
Q3: true
Q4: false
Q5: false
Q6: false
L1: true
L2: false
L3: false
L4: true
L5: false

Q1: false
Q2: false
Q3: false
Q4: true
Q5: false
Q6: false
L1: false
L2: false
L3: true
L4: false
L5: false

Q1: true
Q2: false
Q3: false
Q4: true
Q5: false
Q6: false
L1: true
L2: false
L3: true
L4: false
L5: false

Q1: false
Q2: true
Q3: false
Q4: true
Q5: false
Q6: false
L1: false
L2: false
L3: true
L4: false
L5: false

Q1: true
Q2: true
Q3: false
Q4: true
Q5: false
Q6: false
L1: true
L2: false
L3: true
L4: false
L5: false

Q1: false
Q2: false
Q3: true
Q4: true
Q5: false
Q6: false
L1: false
L2: false
L3: true
L4: false
L5: false

Q1: true
Q2: false
Q3: true
Q4: true
Q5: false
Q6: false
L1: true
L2: false
L3: true
L4: false
L5: false

Q1: false
Q2: true
Q3: true
Q4: true
Q5: false
Q6: false
L1: false
L2: false
L3: true
L4: false
L5: false

Q1: true
Q2: true
Q3: true
Q4: true
Q5: false
Q6: false
L1: true
L2: false
L3: true
L4: false
L5: false

Q1: false
Q2: false
Q3: false
Q4: false
Q5: true
Q6: false
L1: false
L2: false
L3: false
L4: true
L5: true

Q1: true
Q2: false
Q3: false
Q4: false
Q5: true
Q6: false
L1: true
L2: false
L3: false
L4: true
L5: true

Q1: false
Q2: true
Q3: false
Q4: false
Q5: true
Q6: false
L1: false
L2: false
L3: false
L4: true
L5: true

Q1: true
Q2: true
Q3: false
Q4: false
Q5: true
Q6: false
L1: true
L2: false
L3: false
L4: true
L5: true

Q1: false
Q2: false
Q3: true
Q4: false
Q5: true
Q6: false
L1: false
L2: false
L3: false
L4: true
L5: true

Q1: true
Q2: false
Q3: true
Q4: false
Q5: true
Q6: false
L1: true
L2: false
L3: false
L4: true
L5: true

Q1: false
Q2: true
Q3: true
Q4: false
Q5: true
Q6: false
L1: false
L2: false
L3: false
L4: true
L5: true

Q1: true
Q2: true
Q3: true
Q4: false
Q5: true
Q6: false
L1: true
L2: false
L3: false
L4: true
L5: true

Q1: false
Q2: false
Q3: false
Q4: true
Q5: true
Q6: false
L1: false
L2: false
L3: false
L4: true
L5: true

Q1: true
Q2: false
Q3: false
Q4: true
Q5: true
Q6: false
L1: true
L2: false
L3: false
L4: true
L5: true

Q1: false
Q2: true
Q3: false
Q4: true
Q5: true
Q6: false
L1: false
L2: false
L3: false
L4: true
L5: true

Q1: true
Q2: true
Q3: false
Q4: true
Q5: true
Q6: false
L1: true
L2: false
L3: false
L4: true
L5: true

Q1: false
Q2: false
Q3: true
Q4: true
Q5: true
Q6: false
L1: false
L2: false
L3: false
L4: true
L5: true

Q1: true
Q2: false
Q3: true
Q4: true
Q5: true
Q6: false
L1: true
L2: false
L3: false
L4: true
L5: true

Q1: false
Q2: true
Q3: true
Q4: true
Q5: true
Q6: false
L1: false
L2: false
L3: false
L4: true
L5: true

Q1: true
Q2: true
Q3: true
Q4: true
Q5: true
Q6: false
L1: true
L2: false
L3: false
L4: true
L5: true

Q1: false
Q2: false
Q3: false
Q4: false
Q5: false
Q6: true
L1: false
L2: false
L3: false
L4: true
L5: false

Q1: true
Q2: false
Q3: false
Q4: false
Q5: false
Q6: true
L1: true
L2: false
L3: false
L4: true
L5: false

Q1: false
Q2: true
Q3: false
Q4: false
Q5: false
Q6: true
L1: false
L2: false
L3: false
L4: true
L5: false

Q1: true
Q2: true
Q3: false
Q4: false
Q5: false
Q6: true
L1: true
L2: false
L3: false
L4: true
L5: false

Q1: false
Q2: false
Q3: true
Q4: false
Q5: false
Q6: true
L1: false
L2: false
L3: false
L4: true
L5: false

Q1: true
Q2: false
Q3: true
Q4: false
Q5: false
Q6: true
L1: true
L2: false
L3: false
L4: true
L5: false

Q1: false
Q2: true
Q3: true
Q4: false
Q5: false
Q6: true
L1: false
L2: false
L3: false
L4: true
L5: false

Q1: true
Q2: true
Q3: true
Q4: false
Q5: false
Q6: true
L1: true
L2: false
L3: false
L4: true
L5: false

Q1: false
Q2: false
Q3: false
Q4: true
Q5: false
Q6: true
L1: false
L2: false
L3: true
L4: false
L5: false

Q1: true
Q2: false
Q3: false
Q4: true
Q5: false
Q6: true
L1: true
L2: false
L3: true
L4: false
L5: false

Q1: false
Q2: true
Q3: false
Q4: true
Q5: false
Q6: true
L1: false
L2: false
L3: true
L4: false
L5: false

Q1: true
Q2: true
Q3: false
Q4: true
Q5: false
Q6: true
L1: true
L2: false
L3: true
L4: false
L5: false

Q1: false
Q2: false
Q3: true
Q4: true
Q5: false
Q6: true
L1: false
L2: false
L3: true
L4: false
L5: false

Q1: true
Q2: false
Q3: true
Q4: true
Q5: false
Q6: true
L1: true
L2: false
L3: true
L4: false
L5: false

Q1: false
Q2: true
Q3: true
Q4: true
Q5: false
Q6: true
L1: false
L2: false
L3: true
L4: false
L5: false

Q1: true
Q2: true
Q3: true
Q4: true
Q5: false
Q6: true
L1: true
L2: false
L3: true
L4: false
L5: false

Q1: false
Q2: false
Q3: false
Q4: false
Q5: true
Q6: true
L1: false
L2: false
L3: false
L4: true
L5: true

Q1: true
Q2: false
Q3: false
Q4: false
Q5: true
Q6: true
L1: true
L2: false
L3: false
L4: true
L5: true

Q1: false
Q2: true
Q3: false
Q4: false
Q5: true
Q6: true
L1: false
L2: false
L3: false
L4: true
L5: true

Q1: true
Q2: true
Q3: false
Q4: false
Q5: true
Q6: true
L1: true
L2: false
L3: false
L4: true
L5: true

Q1: false
Q2: false
Q3: true
Q4: false
Q5: true
Q6: true
L1: false
L2: false
L3: false
L4: true
L5: true

Q1: true
Q2: false
Q3: true
Q4: false
Q5: true
Q6: true
L1: true
L2: false
L3: false
L4: true
L5: true

Q1: false
Q2: true
Q3: true
Q4: false
Q5: true
Q6: true
L1: false
L2: false
L3: false
L4: true
L5: true

Q1: true
Q2: true
Q3: true
Q4: false
Q5: true
Q6: true
L1: true
L2: false
L3: false
L4: true
L5: true

Q1: false
Q2: false
Q3: false
Q4: true
Q5: true
Q6: true
L1: false
L2: false
L3: false
L4: true
L5: true

Q1: true
Q2: false
Q3: false
Q4: true
Q5: true
Q6: true
L1: true
L2: false
L3: false
L4: true
L5: true

Q1: false
Q2: true
Q3: false
Q4: true
Q5: true
Q6: true
L1: false
L2: false
L3: false
L4: true
L5: true

Q1: true
Q2: true
Q3: false
Q4: true
Q5: true
Q6: true
L1: true
L2: false
L3: false
L4: true
L5: true

Q1: false
Q2: false
Q3: true
Q4: true
Q5: true
Q6: true
L1: false
L2: false
L3: false
L4: true
L5: true

Q1: true
Q2: false
Q3: true
Q4: true
Q5: true
Q6: true
L1: true
L2: false
L3: false
L4: true
L5: true

Q1: false
Q2: true
Q3: true
Q4: true
Q5: true
Q6: true
L1: false
L2: false
L3: false
L4: true
L5: true

Q1: true
Q2: true
Q3: true
Q4: true
Q5: true
Q6: true
L1: true
L2: false
L3: false
L4: true
L5: true
```

## Quellcode

nandu/nandu.go
```go
package nandu

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Nandu struct {
	Heads   []*Node
	Inputs  []*Node
	Outputs []*Node
	Blocks  map[string]*Node
}

var byteToOperation = map[byte]Operation{
	'W': WhiteBlock,
	'B': BlueBlock,
	'R': RedBlockSensor,
	'r': RedBlockNonSensor,
	'X': AlwaysFalse,
	'Q': Passthrough,
	'L': Passthrough,
}

func Parse(r io.Reader) (nandu *Nandu, err error) {
	scanner := bufio.NewScanner(r)

	// skip over width and height
	scanner.Scan()

	nandu = &Nandu{
		Blocks: make(map[string]*Node),
	}

	var lastId int

	var lastRow []*Node
	var openPartner *Node

	for scanner.Scan() {
		line := scanner.Text()

		var currentRow []*Node
		for x, c := range strings.Fields(line) {
			newNode := NewNode(byteToOperation[c[0]], c)

			switch c[0] {
			case 'Q':
				nandu.Inputs = append(nandu.Inputs, newNode)
			case 'L':
				nandu.Outputs = append(nandu.Outputs, newNode)
			case 'W', 'B', 'R', 'r':
				if openPartner != nil {
					newNode.Partner = openPartner
					openPartner.Partner = newNode

					openPartner = nil
				} else {
					openPartner = newNode
				}
			}
			currentRow = append(currentRow, newNode)

			if lastRow != nil {
				lastRow[x].Next = newNode
			}

			if len(newNode.Identifier) == 1 {
				lastId++
				newNode.Identifier += strconv.Itoa(lastId)
			}

			nandu.Blocks[newNode.Identifier] = newNode
		}

		lastRow = currentRow

		if nandu.Heads == nil {
			nandu.Heads = currentRow
		}
	}

	for _, input := range nandu.Inputs {
		input.Set(false)
	}

	return nandu, nil
}
```

nandu/node.go
```go
package nandu

type Operation func(i1, i2 bool) bool

type Node struct {
	Operation   Operation
	I, O        bool
	Next        *Node
	Partner     *Node
	Identifier  string
	initialized bool
}

func NewNode(operator Operation, identifier string) *Node {
	return &Node{
		Operation:  operator,
		Identifier: identifier,
	}
}

func (n *Node) Set(i bool) {
	if n.I == i && n.initialized {
		return
	}
	n.initialized = true

	n.I = i
	n.Update()
	n.updatePartner()
}

func (n *Node) updatePartner() {
	if n.Partner != nil {
		n.Partner.Update()
	}
}

func (n *Node) Update() {
	n.O = n.Operation(n.I, n.Partner != nil && n.Partner.I)
	if n.Next != nil {
		n.Next.Set(n.O)
	}
}
```

nandu/operations.go
```go
package nandu

func BlueBlock(i1, _ bool) bool {
	return i1
}

func RedBlockNonSensor(_, i2 bool) bool {
	return !i2
}

func RedBlockSensor(i1, _ bool) bool {
	return !i1
}

func WhiteBlock(i1, i2 bool) bool {
	o := !(i1 && i2)
	return o
}

func Passthrough(i1, _ bool) bool {
	return i1
}

func AlwaysFalse(_, _ bool) bool {
	return false
}
```

nandu/main.go
```go
package main

import (
	"BWINF/nandu/nandu"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: nandu <file>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	n, err := nandu.Parse(file)
	if err != nil {
		panic(err)
	}

	for x := 0; x < 1<<len(n.Inputs); x++ {
		for i, input := range n.Inputs {
			value := (x>>i)&1 == 1
			fmt.Printf("%v: %v\n", input.Identifier, value)

			input.Set(value)
		}

		for _, output := range n.Outputs {
			fmt.Printf("%v: %v\n", output.Identifier, output.O)
		}

		fmt.Println()
	}
}
```