package main

import "math"

type vec3 struct {
	x, y, z float64
}

func (v vec3) add(u vec3) vec3 {
	return vec3{
		x: v.x + u.x,
		y: v.y + u.y,
		z: v.z + u.z,
	}
}

func (v vec3) sub(u vec3) vec3 {
	return v.add(u.scale(-1))
}

func (v vec3) cross(u vec3) vec3 {
	return vec3{
		x: v.y*u.z - v.z*u.y,
		y: v.z*u.x - v.x*u.z,
		z: v.x*u.y - v.y*u.x,
	}
}

func (v vec3) unit() vec3 {
	return v.scale(1 / v.norm())
}

func (v vec3) scale(f float64) vec3 {
	return vec3{
		x: v.x * f,
		y: v.y * f,
		z: v.z * f,
	}
}

func (v vec3) norm() float64 {
	return math.Sqrt(v.norm2())
}

func (v vec3) norm2() float64 {
	return v.dot(v)
}

func (v vec3) dot(u vec3) float64 {
	return v.x*u.x + v.y*u.y + v.z*u.z
}

func (v vec3) max(u vec3) vec3 {
	return vec3{
		x: math.Max(v.x, u.x),
		y: math.Max(v.y, u.y),
		z: math.Max(v.z, u.z),
	}
}

func (v vec3) min(u vec3) vec3 {
	return vec3{
		x: math.Min(v.x, u.x),
		y: math.Min(v.y, u.y),
		z: math.Min(v.z, u.z),
	}
}
