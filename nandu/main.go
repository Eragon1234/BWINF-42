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
			fmt.Printf("Q%v: %v\n", i+1, value)

			input.Set(value)
		}

		for i, output := range n.Outputs {
			fmt.Printf("L%v: %v\n", i+1, output.Value)
		}

		fmt.Println()
	}
}
