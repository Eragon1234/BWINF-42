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

	for _, input := range n.Inputs {
		input.Set(true)
	}

	for _, output := range n.Outputs {
		fmt.Println(output.O)
	}
}
