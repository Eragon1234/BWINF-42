package main

import (
	"BWINF/zauberschule/floorplan"
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

	plan, err := floorplan.Parse(file)
	if err != nil {
		panic(err)
	}

	path := floorplan.FindPath(plan)

	fmt.Println(path)
}
