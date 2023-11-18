package nandu

import (
	"bufio"
	"io"
	"strings"
)

type Nandu struct {
	Heads   []*Node
	Inputs  []*Node
	Outputs []*Node
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

	nandu = &Nandu{}

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
