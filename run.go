package main

import (
	"image/png"
	"math"
	"math/rand"
	"os"
	"path/filepath"
)

func run(pxWide, pxHigh int, filename string, samples int, debug bool) error {
	cam := newCamera(cameraConfig{
		location:    vec3{z: 10},
		lookingAt:   vec3{},
		upDirection: vec3{y: 1},
		fovDegrees:  20,
		focalLength: 3,
		focalRatio:  math.MaxFloat64,
	})

	fn := union(
		sphere(vec3{}, 1),
		box(vec3{-1.5, -2, -1.5}, vec3{+1.5, -1, +1.5}),
	)

	sk := skySum(
		sun(
			vec3{y: 5, x: 1, z: 2},
			10.0,
			fcolor{1, 1, 1},
		),
		baseSky(fcolor{0.0005, 0.0005, 0.0010}),
	)

	acc := newAccumulator(pxWide, pxHigh)
	rng := rand.New(rand.NewSource(0))
	eng := &engine{acc, cam, fn, sk, rng, samples, false}
	eng.renderFrame()
	if err := writeAccAsImage(acc, filename, false); err != nil {
		return err
	}

	if debug {
		acc := newAccumulator(pxWide, pxHigh)
		rng := rand.New(rand.NewSource(0))
		eng := &engine{acc, cam, fn, sk, rng, 1, true}
		eng.renderFrame()
		if err := writeAccAsImage(acc, debugNormsFilename(filename), true); err != nil {
			return err
		}
	}

	return nil
}

func debugNormsFilename(filename string) string {
	ext := filepath.Ext(filename)
	return filename[:len(ext)-1] + "_debug_norms" + ext
}

func writeAccAsImage(acc *accumulator, filename string, raw bool) error {
	img := acc.image(raw)
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	if err := png.Encode(f, img); err != nil {
		return err
	}
	return f.Close()
}
