package main

import (
	"fmt"
	"image"
	"os"
	"strings"

	_ "image/jpeg"
	_ "image/png"
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

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	var sb strings.Builder

	bounds := img.Bounds()

	var x, y int

	for {
		r, g, b := GetRGB(img, x, y)

		x += int(g)
		y += int(b)

		x %= bounds.Max.X
		y %= bounds.Max.Y

		sb.WriteRune(rune(r))

		if g == 0 && b == 0 {
			break
		}
	}

	fmt.Println(sb.String())
}

func GetRGB(img image.Image, x, y int) (r, g, b uint32) {
	r, g, b, _ = img.At(x, y).RGBA()
	r >>= 8
	g >>= 8
	b >>= 8
	return
}
