package main

import (
	"image/color"
)

type fcolor struct {
	rgb vec3
}

func (c fcolor) color() color.RGBA {
	const w = 0x100
	lim := c.rgb.max(vec3{}).min(vec3{1, 1, 1})
	return color.RGBA{
		R: uint8(lim.x * w),
		G: uint8(lim.y * w),
		B: uint8(lim.z * w),
		A: w - 1,
	}
}
