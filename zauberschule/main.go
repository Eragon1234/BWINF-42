package main

import (
	"BWINF/pkg/slice"
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: zauberschule <filepath>")
		return
	}

	filepath := os.Args[1]

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	var n, m int
	_, err = fmt.Sscanf(scanner.Text(), "%d %d", &n, &m)
	if err != nil {
		panic(err)
	}

	plan := [2][][]bool{
		slice.New2D[bool](n, m),
		slice.New2D[bool](n, m),
	}
}
