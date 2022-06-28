package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
)

func main() {
	f, err := os.Open("load-png/image.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	bounds := img.Bounds().Max
	fmt.Printf("Image size: %dpx x %dpx\n", bounds.X, bounds.Y)
}
