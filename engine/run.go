package engine

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
)

type RunArgs struct {
}

func Run(args RunArgs) ([]byte, error) {
	var buf bytes.Buffer
	if err := gif.EncodeAll(&buf, dummyGIF()); err != nil {
		return nil, fmt.Errorf("encoding gif: %w", err)
	}
	return buf.Bytes(), nil
}

func dummyGIF() *gif.GIF {
	const (
		pxWide = 160
		pxHigh = 128
	)
	solid := func(c color.Color) *image.Paletted {
		return &image.Paletted{
			Pix:    make([]uint8, pxWide*pxHigh),
			Stride: pxWide,
			Rect: image.Rectangle{
				Min: image.Point{X: 0, Y: 0},
				Max: image.Point{X: pxWide, Y: pxHigh},
			},
			Palette: []color.Color{c},
		}
	}
	imgs := []*image.Paletted{
		solid(color.RGBA{0xFF, 0x00, 0x00, 0xFF}),
		solid(color.RGBA{0x00, 0xFF, 0x00, 0xFF}),
		solid(color.RGBA{0x00, 0x00, 0xFF, 0xFF}),
	}
	return &gif.GIF{
		Image: imgs,
		Delay: []int{4, 4, 4},
	}
}
