package main

type ray struct {
	o, d vec3
}

func (r ray) at(t float64) vec3 {
	return r.o.add(r.d.scale(t))
}
