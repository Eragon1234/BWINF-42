package main

import (
	"BWINF/pkg/ansi"
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
	fmt.Println("Legende:")
	fmt.Println("Der Startpunkt ist", ansi.S("unterstrichen.", ansi.Underline))
	fmt.Println("Der Endpunkt ist X.")
	fmt.Println("Stockwerkwechsel sind farbkodiert.")
	fmt.Println("Bewegungen sind mit Pfeilen dargestellt.")
	fmt.Println()

	fmt.Println("Weg als Liste, Start nicht inkludiert:")
	for _, step := range *shortestPath {
		fmt.Printf("%+v\n", step)
	}

	fmt.Println()

	fmt.Printf("Der kürzeste Weg braucht %d Sekunden.\n", distance)
}
