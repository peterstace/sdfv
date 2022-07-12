package main

import "math/rand"

type engine struct {
	acc     *accumulator
	cam     *camera
	fn      sdf
	sk      sky
	rng     *rand.Rand
	samples int

	debugNorms bool
}

func (e *engine) renderFrame() {
	pxPitch := 2.0 / float64(e.acc.pxWide)
	for pxY := 0; pxY < e.acc.pxHigh; pxY++ {
		y := (float64(pxY-e.acc.pxHigh/2) + e.rng.Float64()) * pxPitch * -1.0
		for pxX := 0; pxX < e.acc.pxWide; pxX++ {
			x := (float64(pxX-e.acc.pxWide/2) + e.rng.Float64()) * pxPitch
			for s := 0; s < e.samples; s++ {
				r := e.cam.makeRay(x, y, e.rng)
				var fc fcolor
				if e.debugNorms {
					fc = e.traceNormal(r)
				} else {
					fc = e.trace(r)
				}
				e.acc.add(pxX, pxY, fc)
			}
		}
	}
}

func (e *engine) trace(r ray) fcolor {
	t, ok := e.findSurface(r)
	if !ok {
		return e.sk(r.uDir)
	}
	hit := r.at(t)

	uNorm := e.uNormal(hit)
	uHemi := vec3{
		e.rng.NormFloat64(),
		e.rng.NormFloat64(),
		e.rng.NormFloat64(),
	}.unit()
	if uHemi.dot(uNorm) < 0 {
		uHemi = uHemi.scale(-1)
	}

	const bumpOut = 0.001
	nextR := ray{
		origin: hit.add(uNorm.scale(bumpOut)),
		uDir:   uHemi,
	}
	brdf := uHemi.dot(uNorm)
	cf := e.trace(nextR)
	cf.rgb = cf.rgb.scale(brdf)
	return cf
}

func (e *engine) traceNormal(r ray) fcolor {
	t, ok := e.findSurface(r)
	if !ok {
		return fcolor{}
	}
	n := e.uNormal(r.at(t))
	return unitDirToColor(n)
}

func unitDirToColor(uDir vec3) fcolor {
	return fcolor{rgb: uDir.add(vec3{1, 1, 1}).scale(0.5)}
}

func (e *engine) uNormal(v vec3) vec3 {
	const eps = 1e-6
	offsetX := e.fn(v.add(vec3{x: eps}))
	offsetY := e.fn(v.add(vec3{y: eps}))
	offsetZ := e.fn(v.add(vec3{z: eps}))
	offset0 := e.fn(v)
	return vec3{offsetX, offsetY, offsetZ}.
		sub(vec3{offset0, offset0, offset0}).
		unit()
}

func (e *engine) findSurface(r ray) (float64, bool) {
	const (
		escapeThreshold    = 100
		intersectThreshold = 0.001
		maxIterations      = 1e8
	)
	t := 0.0
	for i := 0; i < maxIterations; i++ {
		d := e.fn(r.at(t))
		if d >= escapeThreshold {
			return 0, false
		}
		t += d
		if d < intersectThreshold {
			return t, true
		}
	}
	panic("reached max iterations")
}
