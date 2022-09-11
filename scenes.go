package main

import "math"

type scene struct {
	cam *camera
	fn  sdf
	sk  sky
}

func lookupScene(sceneName string) (scene, bool) {
	scn, ok := map[string]scene{
		"original": originalScene,
	}[sceneName]
	return scn, ok
}

var originalScene = scene{
	cam: newCamera(cameraConfig{
		location:    vec3{z: 10},
		lookingAt:   vec3{},
		upDirection: vec3{y: 1},
		fovDegrees:  20,
		focalLength: 3,
		focalRatio:  math.MaxFloat64,
	}),
	fn: union(
		sphere(vec3{}, 1),
		box(vec3{-1.5, -2, -1.5}, vec3{+1.5, -1, +1.5}),
	),
	sk: skySum(
		sun(
			vec3{y: 5, x: 1, z: 2},
			10.0,
			fcolor{1, 1, 1},
		),
		baseSky(fcolor{0.0005, 0.0005, 0.0010}),
	),
}
