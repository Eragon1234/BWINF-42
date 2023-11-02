package main

import (
	"BWINF/zauberschule/floorplan"
	"BWINF/zauberschule/path"
	"BWINF/zauberschule/route"
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

	p, distance := path.FindPath(plan)

	fmt.Println(route.StringFloorplanWithPath(*plan, *p))

	fmt.Println("Distance: ", distance)
}
