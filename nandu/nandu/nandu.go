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
			var n *block.Node
			switch c[0] {
			case 'Q':
				n = block.NewNode(&block.NoOp{}, c)
				nandu.Inputs = append(nandu.Inputs, n)
			case 'L':
				n = block.NewNode(&block.NoOp{}, c)
				nandu.Outputs = append(nandu.Outputs, n)
			case 'W', 'B', 'R', 'r':
				b := block.NewBlock(stringToType[c])
				n = block.NewNode(b, c)
				if openPartner != nil {
					connect(openPartner, n)
					openPartner = nil
				} else {
					openPartner = n
				}
			case 'X':
				n = block.NewNode(&block.NoBlock{}, c)
			}
			currentRow = append(currentRow, n)

			if lastRow != nil {
				lastRow[x].Next = n
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

	n1.Dependencies = append(n1.Dependencies, n2)
	n2.Dependencies = append(n2.Dependencies, n1)
}
