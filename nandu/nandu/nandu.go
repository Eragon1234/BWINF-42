package nandu

import (
	"BWINF/nandu/nandu/block"
	"bufio"
	"io"
	"strings"
)

type Nandu struct {
	Inputs  []*block.Node
	Outputs []*block.Node
}

var stringToType = map[string]func(i1, i2 bool) bool{
	"W": block.WhiteBlock,
	"B": block.BlueBlock,
	"R": block.RedBlockSensor,
	"r": block.RedBlockNonSensor,
}

func Parse(r io.Reader) (nandu *Nandu, err error) {
	scanner := bufio.NewScanner(r)

	// skip over width and height
	scanner.Scan()

	nandu = &Nandu{}

	var openPartner *block.Node
	var currentRow, lastRow []*block.Node

	for scanner.Scan() {
		line := scanner.Text()

		for x, c := range strings.Fields(line) {
			var newNode *block.Node
			switch c[0] {
			case 'Q':
				newNode = block.NewNode(&block.NoOp{}, c)
				nandu.Inputs = append(nandu.Inputs, newNode)
			case 'L':
				newNode = block.NewNode(&block.NoOp{}, c)
				nandu.Outputs = append(nandu.Outputs, newNode)
			case 'W', 'B', 'R', 'r':
				b := block.NewBlock(stringToType[c])
				newNode = block.NewNode(b, c)
				if openPartner != nil {
					connect(openPartner, newNode)
					openPartner = nil
				} else {
					openPartner = newNode
				}
			case 'X':
				newNode = block.NewNode(&block.NoBlock{}, c)
			}
			currentRow = append(currentRow, newNode)

			if lastRow != nil {
				lastRow[x].Next = newNode
			}
		}

		lastRow = currentRow
		currentRow = nil
	}

	return nandu, nil
}

// connect connects two nodes
// It expects that n1 and n2 contain a block
func connect(n1 *block.Node, n2 *block.Node) {
	n1.Op.(*block.Block).Partner = n2
	n2.Op.(*block.Block).Partner = n1

	n1.Partner = n2
	n2.Partner = n1
}
