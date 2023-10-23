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

	var start, end [3]int

	for i := 0; i < n; i++ {
		scanner.Scan()
		for j, c := range scanner.Text() {
			if c == 'A' {
				start = [3]int{0, i, j}
			} else if c == 'B' {
				end = [3]int{0, i, j}
			}
			plan[0][i][j] = c != '#'
		}
	}

	// skip the empty line
	scanner.Scan()

	for i := 0; i < n; i++ {
		scanner.Scan()
		for j, c := range scanner.Text() {
			if c == 'A' {
				start = [3]int{1, i, j}
			} else if c == 'B' {
				end = [3]int{1, i, j}
			}

			plan[1][i][j] = c != '#'
		}
	}

	fmt.Println(start, end)

	for _, s := range plan {
		for _, row := range s {
			for _, c := range row {
				if c {
					fmt.Print(".")
				} else {
					fmt.Print("#")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
