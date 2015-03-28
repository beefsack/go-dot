package dot

import (
	"image"
	"image/color"
)

func FromImage(img image.Image, width, height int, p Pixeler) [][]bool {
	pix := make([][]bool, height)
	bounds := img.Bounds()
	xRatio := float64(bounds.Dx()) / float64(width)
	yRatio := float64(bounds.Dy()) / float64(height)
	perX := int(xRatio)
	perY := int(yRatio)

	for y := 0; y < height; y++ {
		pix[y] = make([]bool, width)
		for x := 0; x < width; x++ {
			p0 := image.Point{
				bounds.Min.X + int(float64(x)*xRatio),
				bounds.Min.Y + int(float64(y)*yRatio),
			}
			pix[y][x] = p(img, image.Rectangle{
				p0,
				image.Point{
					p0.X + perX,
					p0.Y + perY,
				},
			})
		}
	}
	return pix
}

type Pixeler func(img image.Image, bounds image.Rectangle) bool

type ColorFilterer func(c color.Color) bool

func MostCommonColor(img image.Image, bounds image.Rectangle) color.Color {
	max := color.Color(color.RGBA{})
	maxCount := 0
	counts := map[color.Color]int{}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			counts[c] += 1
			if counts[c] > maxCount {
				max = c
				maxCount = counts[c]
			}
		}
	}
	return max
}

func AverageColor(img image.Image, bounds image.Rectangle) color.Color {
	num := uint64(bounds.Dx() * bounds.Dy())
	if num == 0 {
		return color.RGBA{}
	}
	var r, g, b, a uint64
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			cr, cg, cb, ca := c.RGBA()
			r += uint64(cr)
			g += uint64(cg)
			b += uint64(cb)
			a += uint64(ca)
		}
	}
	return color.RGBA{
		uint8(r / num),
		uint8(g / num),
		uint8(b / num),
		uint8(a / num),
	}
}

func ColorFilterPerc(
	img image.Image,
	bounds image.Rectangle,
	filter ColorFilterer,
) float64 {
	total := uint64(bounds.Dx() * bounds.Dy())
	if total == 0 {
		return 0
	}
	num := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if filter(img.At(x, y)) {
				num++
			}
		}
	}
	return float64(num) / float64(total)
}

func Saturation(c color.Color) uint8 {
	r, g, b, a := c.RGBA()
	return uint8((r + g + b) / 3 * a / 256)
}
