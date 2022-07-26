package main

import (
	"image"
)

type accumulator struct {
	pxWide int
	pxHigh int
	data   []fcolor
}

func newAccumulator(pxWide, pxHigh int) *accumulator {
	return &accumulator{
		pxWide: pxWide,
		pxHigh: pxHigh,
		data:   make([]fcolor, pxWide*pxHigh),
	}
}

func (a *accumulator) add(pxX, pxY int, c fcolor) {
	idx := a.idx(pxX, pxY)
	a.data[idx] = a.data[idx].add(c)
}

func (a *accumulator) get(pxX, pxY int) fcolor {
	idx := a.idx(pxX, pxY)
	return a.data[idx]
}

func (a *accumulator) idx(pxX, pxY int) int {
	return pxX + pxY*a.pxWide
}

func (a *accumulator) image(raw bool) image.Image {
	const (
		gamma    = 2.2
		exposure = 1.0
	)
	mean := a.mean()
	img := image.NewRGBA(image.Rectangle{
		Max: image.Pt(a.pxWide, a.pxHigh),
	})
	for pxY := 0; pxY < a.pxHigh; pxY++ {
		for pxX := 0; pxX < a.pxWide; pxX++ {
			fc := a.get(pxX, pxY)
			if !raw {
				fc = fc.scale(0.5 * exposure / mean).pow(1 / gamma)
			}
			img.Set(pxX, pxY, fc.color())
		}
	}
	return img
}

func (a *accumulator) mean() float64 {
	var sum float64
	for _, c := range a.data {
		sum += c.r + c.g + c.b
	}
	return sum / float64(len(a.data)) / 3.0
}
