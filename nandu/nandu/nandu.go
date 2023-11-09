package nandu

import (
	"BWINF/nandu/nandu/block"
	"bufio"
	"io"
	"strings"
)

type Nandu struct {
	Inputs  []*block.Input
	Outputs []*block.Output
}

func Parse(r io.Reader) (nandu *Nandu, err error) {
	scanner := bufio.NewScanner(r)

	// skip over width and height
	scanner.Scan()

	nandu = &Nandu{}

	var openPartner *block.Block
	var currentRow, lastRow []block.Node

	for scanner.Scan() {
		line := scanner.Text()

		for x, c := range strings.Fields(line) {
			var b block.Node
			switch c[0] {
			case 'Q':
				b = &block.Input{}
				nandu.Inputs = append(nandu.Inputs, b.(*block.Input))
			case 'L':
				b = &block.Output{}
				nandu.Outputs = append(nandu.Outputs, b.(*block.Output))
			case 'W':
				b = block.NewBlock(block.WhiteBlock)
			case 'B':
				b = block.NewBlock(block.BlueBlock)
			case 'R':
				b = block.NewBlock(block.RedBlockSensor)
			case 'r':
				b = block.NewBlock(block.RedBlockNonSensor)
			case 'X':
				b = new(block.NoBlock)
			}
			currentRow = append(currentRow, b)

			if lastRow != nil {
				lastRow[x].SetNext(b)
			}

			if b, ok := b.(*block.Block); ok {
				if openPartner != nil {
					openPartner.SetPartner(b)
					openPartner = nil
				} else {
					openPartner = b
				}
			}
		}

		lastRow = currentRow
		currentRow = nil
	}

	return nandu, nil
}
