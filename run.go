package main

import (
	"image"
	"image/png"
	"math"
	"math/rand"
	"os"
)

func run(pxWide, pxHigh int, filename string) error {
	cam := newCamera(cameraConfig{
		location:    vec3{x: -1, y: +2, z: +3},
		lookingAt:   vec3{},
		upDirection: vec3{y: 1},
		fovDegrees:  90,
		focalLength: 3,
		focalRatio:  math.MaxFloat64,
	})

	img := renderImage(pxWide, pxHigh, cam)
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	if err := png.Encode(f, img); err != nil {
		return err
	}
	return f.Close()
}

func renderImage(pxWide, pxHigh int, cam *camera) image.Image {
	rng := rand.New(rand.NewSource(0))
	img := image.NewRGBA(image.Rectangle{
		Max: image.Pt(pxWide, pxHigh),
	})
	pxPitch := 2.0 / float64(pxWide)
	for pxY := 0; pxY < pxHigh; pxY++ {
		y := (float64(pxY-pxHigh/2) + rng.Float64()) * pxPitch * -1.0
		for pxX := 0; pxX < pxWide; pxX++ {
			x := (float64(pxX-pxWide/2) + rng.Float64()) * pxPitch
			r := cam.makeRay(x, y, rng)
			fc := trace(r)
			img.SetRGBA(pxX, pxY, fc.color())
		}
	}
	return img
}

func trace(r ray) fcolor {
	// TODO
	return fcolor{
		rgb: vec3{
			x: rand.Float64(),
			y: rand.Float64(),
			z: rand.Float64(),
		},
	}
}
