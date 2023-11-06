package nandu

import (
	"BWINF/nandu/nandu/block"
	"BWINF/pkg/slice"
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Nandu struct {
	Blocks  [][]*block.Block
	Inputs  []*block.Block
	Outputs []*block.Block
}

func Parse(r io.Reader) (nandu *Nandu, err error) {
	scanner := bufio.NewScanner(r)

	var width, height int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d %d", &width, &height)

	nandu = &Nandu{
		Blocks: slice.New2D[*block.Block](height, width),
	}

	var openPartner *block.Block

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()

		for x, c := range strings.Fields(line) {
			y, x, c := y, x, c

			b := block.NewNoBlock()
			nandu.Blocks[y][x] = b
			b.SetOutput = func(o bool) {
				if y+1 < len(nandu.Blocks) && nandu.Blocks[y+1][x] != nil {
					nandu.Blocks[y+1][x].Set(o)
				}
			}

			switch c {
			case "W":
				b.SetOperation(block.WhiteBlock)
			case "B":
				b.SetOperation(block.BlueBlock)
			case "R":
				b.SetOperation(block.RedBlockSensor)
			case "r":
				b.SetOperation(block.RedBlockNonSensor)
			default:
				if c[0] == 'Q' {
					b.SetOperation(block.PassThrough)
					nandu.Inputs = append(nandu.Inputs, b)
				} else if c[0] == 'L' {
					b.SetOperation(block.PassThrough)
					nandu.Outputs = append(nandu.Outputs, b)
				}
			}

			if c == "X" || c[0] == 'Q' || c[0] == 'L' {
				continue
			}

			if openPartner != nil {
				b.SetPartner(openPartner)
				openPartner = nil
			} else {
				openPartner = b
			}
		}
	}

	return nandu, nil
}
