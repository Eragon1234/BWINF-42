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

	fmt.Println(plan.Start, plan.End)

	for _, s := range plan.Plan {
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
