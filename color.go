package main

import (
	"image/color"
	"math"
)

type fcolor struct {
	r, g, b float64
}

func (c fcolor) add(o fcolor) fcolor {
	return fcolor{
		c.r + o.r,
		c.g + o.g,
		c.b + o.b,
	}
}

func (c fcolor) scale(f float64) fcolor {
	return fcolor{
		c.r * f,
		c.g * f,
		c.b * f,
	}
}

func (c fcolor) pow(exp float64) fcolor {
	return fcolor{
		math.Pow(c.r, exp),
		math.Pow(c.g, exp),
		math.Pow(c.b, exp),
	}
}

func (c fcolor) max(o fcolor) fcolor {
	return fcolor{
		math.Max(c.r, o.r),
		math.Max(c.g, o.g),
		math.Max(c.b, o.b),
	}
}

func (c fcolor) min(o fcolor) fcolor {
	return fcolor{
		math.Min(c.r, o.r),
		math.Min(c.g, o.g),
		math.Min(c.b, o.b),
	}
}

func (c fcolor) color() color.RGBA {
	const w = 0x100
	const max = 1.0 - 0.5/w
	lim := c.max(fcolor{}).min(fcolor{max, max, max})
	return color.RGBA{
		R: uint8(lim.r * w),
		G: uint8(lim.g * w),
		B: uint8(lim.b * w),
		A: w - 1,
	}
}
