package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"

	"github.com/beefsack/go-dot"
)

func main() {
	var (
		r      io.Reader
		file   string
		width  int
		height int
	)
	flag.StringVar(&file, "f", "", "the file to open, otherwise STDIN is used")
	flag.IntVar(&width, "w", 160, "the width in dots, characters is 2 dots wide")
	flag.IntVar(&height, "h", 0, "the height in characters, characters are 4 dots high, defaults to preserving aspect ratio with width")
	flag.Parse()

	if file != "" {
		f, err := os.Open(file)
		if err != nil {
			log.Fatalf("failed to open file, %s", err)
		}
		defer f.Close()
		r = f
	} else {
		r = os.Stdin
	}

	img, _, err := image.Decode(r)
	if err != nil {
		log.Fatalf("failed to decode image, %s", err)
	}

	bounds := img.Bounds()
	if height == 0 && bounds.Dy() > 0 {
		height = width * bounds.Dy() / bounds.Dx()
	}

	fmt.Println(dot.Render(dot.FromImage(
		img,
		width,
		height,
		func(img image.Image, bounds image.Rectangle) bool {
			return dot.Saturation(dot.AverageColor(img, bounds)) > 50
		},
	)))
}
