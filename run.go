package main

import (
	"image/png"
	"math"
	"math/rand"
	"os"
)

func run(pxWide, pxHigh int, filename string) error {
	acc := newAccumulator(pxWide, pxHigh)
	cam := newCamera(cameraConfig{
		location:    vec3{z: 10},
		lookingAt:   vec3{},
		upDirection: vec3{y: 1},
		fovDegrees:  20,
		focalLength: 3,
		focalRatio:  math.MaxFloat64,
	})
	fn := sphere
	sk := skySum(
		sun(
			vec3{y: 5, x: 1, z: 2},
			10.0,
			fcolor{vec3{1, 1, 1}},
		),
		baseSky(fcolor{vec3{0.1, 0.1, 0.2}}),
	)
	rng := rand.New(rand.NewSource(0))
	eng := &engine{acc, cam, fn, sk, rng}
	eng.renderFrame()

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
