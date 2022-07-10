package main

type ray struct {
	origin vec3 // origin
	uDir   vec3 // direction (unit vector)
}

func (r ray) at(t float64) vec3 {
	return r.origin.add(r.uDir.scale(t))
}
