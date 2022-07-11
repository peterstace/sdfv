package main

import "math"

type sdf func(vec3) float64

func sphere(c vec3, r float64) sdf {
	return func(v vec3) float64 {
		return v.sub(c).norm() - r
	}
}

func lowerX(x float64) sdf {
	return func(v vec3) float64 {
		return v.x - x
	}
}

func upperX(x float64) sdf {
	return func(v vec3) float64 {
		return x - v.x
	}
}

func lowerY(y float64) sdf {
	return func(v vec3) float64 {
		return v.y - y
	}
}

func upperY(y float64) sdf {
	return func(v vec3) float64 {
		return y - v.y
	}
}

func lowerZ(z float64) sdf {
	return func(v vec3) float64 {
		return v.z - z
	}
}

func upperZ(z float64) sdf {
	return func(v vec3) float64 {
		return z - v.z
	}
}

func union(fn1, fn2 sdf) sdf {
	return func(v vec3) float64 {
		d1 := fn1(v)
		d2 := fn2(v)
		return math.Min(d1, d2)
	}
}

func intersection(fn1, fn2 sdf) sdf {
	return func(v vec3) float64 {
		d1 := fn1(v)
		d2 := fn2(v)
		return math.Max(d1, d2)
	}
}

func box(min, max vec3) sdf {
	return func(v vec3) float64 {
		t0 := lowerX(max.x)(v)
		t1 := upperX(min.x)(v)
		t2 := lowerY(max.y)(v)
		t3 := upperY(min.y)(v)
		t4 := lowerZ(max.z)(v)
		t5 := upperZ(min.z)(v)
		t := math.Max(t0, t1)
		t = math.Max(t, t2)
		t = math.Max(t, t3)
		t = math.Max(t, t4)
		t = math.Max(t, t5)
		return t
	}
}
