package main

import (
	"image"
)

type accumulator struct {
	pxWide int
	pxHigh int
	data   []fcolor
}

func newAccumulator(pxWide, pxHigh int) accumulator {
	return accumulator{
		pxWide: pxWide,
		pxHigh: pxHigh,
		data:   make([]fcolor, pxWide*pxHigh),
	}
}

func (a accumulator) add(pxX, pxY int, c fcolor) {
	idx := a.idx(pxX, pxY)
	a.data[idx] = fcolor{a.data[idx].rgb.add(c.rgb)}
}

func (a accumulator) get(pxX, pxY int) fcolor {
	idx := a.idx(pxX, pxY)
	return a.data[idx]
}

func (a accumulator) idx(pxX, pxY int) int {
	return pxX + pxY*a.pxWide
}

func (a accumulator) image() image.Image {
	img := image.NewRGBA(image.Rectangle{
		Max: image.Pt(a.pxWide, a.pxHigh),
	})
	for pxY := 0; pxY < a.pxHigh; pxY++ {
		for pxX := 0; pxX < a.pxWide; pxX++ {
			fc := a.get(pxX, pxY)
			img.Set(pxX, pxY, fc.color())
		}
	}
	return img
}
