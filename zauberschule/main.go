package main

import (
	"BWINF/zauberschule/path"
	"BWINF/zauberschule/plan"
	"BWINF/zauberschule/route"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: zauberschule <filepath>")
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

	p, err := plan.Parse(file)
	if err != nil {
		panic(err)
	}

	shortestPath, distance := path.FindPath(p)

	fmt.Println(route.StringFloorplanWithPath(*p, *shortestPath))

	fmt.Println("Distance: ", distance)
}
