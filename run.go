package main

import (
	"image/png"
	"math"
	"math/rand"
	"os"
)

func run(pxWide, pxHigh int, filename string) error {
	cam := newCamera(cameraConfig{
		location:    vec3{z: 10},
		lookingAt:   vec3{},
		upDirection: vec3{y: 1},
		fovDegrees:  20,
		focalLength: 3,
		focalRatio:  math.MaxFloat64,
	})

	acc := renderFrame(pxWide, pxHigh, cam)
	img := acc.image()
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	if err := png.Encode(f, img); err != nil {
		return err
	}
	return f.Close()
}

func renderFrame(pxWide, pxHigh int, cam *camera) accumulator {
	rng := rand.New(rand.NewSource(0))
	acc := newAccumulator(pxWide, pxHigh)
	pxPitch := 2.0 / float64(pxWide)
	for pxY := 0; pxY < pxHigh; pxY++ {
		y := (float64(pxY-pxHigh/2) + rng.Float64()) * pxPitch * -1.0
		for pxX := 0; pxX < pxWide; pxX++ {
			x := (float64(pxX-pxWide/2) + rng.Float64()) * pxPitch
			r := cam.makeRay(x, y, rng)

			sk := skySum(
				sun(
					vec3{y: 5, x: 1, z: 2},
					10.0,
					fcolor{vec3{1, 1, 1}},
				),
				baseSky(fcolor{vec3{0.1, 0.1, 0.2}}),
			)
			fc := trace(r, sphere, sk, rng)

			acc.add(pxX, pxY, fc)
		}
	}
	return acc
}

func trace(r ray, fn sdf, sk sky, rng *rand.Rand) fcolor {
	t, ok := findSurface(r, fn)
	if !ok {
		return sk(r.uDir)
	}
	hit := r.at(t)

	uNorm := uNormal(hit, fn)
	uHemi := vec3{
		rng.NormFloat64(),
		rng.NormFloat64(),
		rng.NormFloat64(),
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
	cf := trace(nextR, fn, sk, rng)
	cf.rgb = cf.rgb.scale(brdf)
	return cf
}

func unitDirToColor(uDir vec3) fcolor {
	return fcolor{rgb: uDir.add(vec3{1, 1, 1}).scale(0.5)}
}

func traceNormal(r ray, fn sdf) fcolor {
	t, ok := findSurface(r, fn)
	if !ok {
		return fcolor{}
	}
	n := uNormal(r.at(t), fn)
	return unitDirToColor(n)
}

func uNormal(v vec3, fn sdf) vec3 {
	const eps = 1e-6
	offsetX := fn(v.add(vec3{x: eps}))
	offsetY := fn(v.add(vec3{y: eps}))
	offsetZ := fn(v.add(vec3{z: eps}))
	offset0 := fn(v)
	return vec3{offsetX, offsetY, offsetZ}.
		sub(vec3{offset0, offset0, offset0}).
		unit()
}

func findSurface(r ray, fn sdf) (float64, bool) {
	const (
		escapeThreshold    = 100
		intersectThreshold = 0.001
		maxIterations      = 1e6
	)
	t := 0.0
	for i := 0; i < maxIterations; i++ {
		d := fn(r.at(t))
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
