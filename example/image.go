package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"

	"github.com/beefsack/go-dot"
)

func main() {
	var (
		r         io.Reader
		file      string
		width     int
		height    int
		threshold float64
	)
	flag.StringVar(&file, "f", "", "the file to open, otherwise STDIN is used")
	flag.IntVar(&width, "w", 160, "the width in dots, characters is 2 dots wide")
	flag.IntVar(&height, "h", 0, "the height in characters, characters are 4 dots high, defaults to preserving aspect ratio with width")
	flag.Float64Var(&threshold, "t", 0.3, "the pixel fill threshold percentage, 0.0-1.0")
	flag.Parse()

	if threshold < 0 || threshold > 1 {
		log.Fatalf("threshold must be between 0.0 and 1.0")
	}

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
			return dot.ColorFilterPerc(img, bounds, func(c color.Color) bool {
				return dot.Saturation(c) > 0
			}) > threshold
		},
	)))
}
