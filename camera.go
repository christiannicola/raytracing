package main

import "math"

type camera struct {
	origin          vec3
	lowerLeftCorner vec3
	horizontal      vec3
	vertical        vec3
}

func newCamera(lookFrom, lookAt, vup vec3, vfov, aspectRatio float64) *camera {
	theta := degreesToRadians(vfov)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := unitVector(subtractVec3(lookFrom, lookAt))
	u := unitVector(cross(vup, w))
	v := cross(w, u)

	horizontal := multiplyVec3ByFactor(u, viewportWidth)
	vertical := multiplyVec3ByFactor(v, viewportHeight)
	lowerLeftCorner := subtractVec3(subtractVec3(subtractVec3(lookFrom, divideVec3(horizontal, 2)), divideVec3(vertical, 2.0)), w)

	return &camera{lookFrom, lowerLeftCorner, horizontal, vertical}
}

func (c *camera) getRay(s, t float64) ray {
	return newRay(c.origin, subtractVec3(addVec3(addVec3(c.lowerLeftCorner, multiplyVec3ByFactor(c.horizontal, s)), multiplyVec3ByFactor(c.vertical, t)), c.origin))
}
