package main

type sdf func(vec3) float64

func sphere(v vec3) float64 {
	return v.norm() - 1
}
