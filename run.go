package main

import (
	"image/png"
	"math"
	"math/rand"
	"os"
)

func run(pxWide, pxHigh int, filename string) error {
	cam := newCamera(cameraConfig{
		location:    vec3{x: 0, y: 0, z: 2},
		lookingAt:   vec3{},
		upDirection: vec3{y: 1},
		fovDegrees:  90,
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
			fc := trace(r, sphere)
			acc.add(pxX, pxY, fc)
		}
	}
	return acc
}

func trace(r ray, fn sdf) fcolor {
	t, ok := distance(r, fn)
	if !ok {
		return fcolor{}
	}
	t *= 0.1
	return fcolor{rgb: vec3{x: t, y: t, z: t}}
}

func distance(r ray, fn sdf) (float64, bool) {
	t := 0.0
	for i := 0; i < 100; i++ {
		d := fn(r.at(t))
		if d >= 100 {
			return 0, false
		}
		t += d
		if d < 0.001 {
			return t, true
		}
	}
	return 0, false
}
