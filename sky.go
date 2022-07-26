package main

import "math"

// sky gives the color of the sky in a given (unit vector) direction.
type sky func(vec3) fcolor

func sun(d vec3, pow float64, col fcolor) sky {
	uDir := d.unit()
	return func(v vec3) fcolor {
		intensity := math.Max(0, v.dot(uDir))
		intensity = math.Pow(intensity, pow)
		return col.scale(intensity)
	}
}

func baseSky(col fcolor) sky {
	return func(vec3) fcolor { return col }
}

func skySum(s1, s2 sky) sky {
	return func(v vec3) fcolor {
		return s1(v).add(s2(v))
	}
}
