package main

import (
	"math"
	"math/rand"
)

type cameraConfig struct {
	location    vec3
	lookingAt   vec3
	upDirection vec3
	fovDegrees  float64
	focalLength float64
	focalRatio  float64
}

type camera struct {
	screenLoc vec3
	screenX   vec3
	screenY   vec3
	eyeLoc    vec3
	eyeX      vec3
	eyeY      vec3
}

func newCamera(cfg cameraConfig) *camera {
	cam := new(camera)

	upDirection := cfg.upDirection.unit()
	viewDirection := cfg.lookingAt.sub(cfg.location).unit()

	cam.screenX = viewDirection.cross(upDirection)
	cam.screenY = cam.screenX.cross(viewDirection)

	cam.eyeX = cam.screenX.scale(cfg.focalLength / cfg.focalRatio)
	cam.eyeY = cam.screenY.scale(cfg.focalLength / cfg.focalRatio)
	cam.eyeLoc = cfg.location

	fovRads := cfg.fovDegrees / 180 * math.Pi
	halfScreenWidth := math.Tan(fovRads/2) * cfg.focalLength
	cam.screenX = cam.screenX.scale(halfScreenWidth)
	cam.screenY = cam.screenY.scale(halfScreenWidth)
	cam.screenLoc = cam.eyeLoc.add(viewDirection.scale(cfg.focalLength))

	return cam
}

func (c *camera) makeRay(x, y float64, rng *rand.Rand) ray {
	start := c.eyeLoc.
		add(c.eyeX.scale(2*rng.Float64() - 1.0)).
		add(c.eyeY.scale(2*rng.Float64() - 1.0))
	end := c.screenLoc.
		add(c.screenX.scale(x)).
		add(c.screenY.scale(y))
	return ray{
		o: start,
		d: end.sub(start).unit(),
	}
}
