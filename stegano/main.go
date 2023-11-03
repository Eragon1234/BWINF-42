package main

import (
	"fmt"
	"image/png"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: stegano <filepath>")
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

	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	var sb strings.Builder

	bounds := img.Bounds()

	var x, y int

	for {
		r, g, b, _ := img.At(x, y).RGBA()
		r >>= 8
		g >>= 8
		b >>= 8
		x += int(g)
		y += int(b)
		for x >= bounds.Max.X {
			x -= bounds.Max.X
		}
		for y >= bounds.Max.Y {
			y -= bounds.Max.Y
		}

		sb.WriteRune(rune(r))

		if g == 0 && b == 0 {
			break
		}
	}

	fmt.Println(sb.String())
}
