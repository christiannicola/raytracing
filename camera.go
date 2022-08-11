package main

import "math"

type camera struct {
	origin          vec3
	lowerLeftCorner vec3
	horizontal      vec3
	vertical        vec3
	u               vec3
	v               vec3
	w               vec3
	lensRadius      float64
}

func newCamera(lookFrom, lookAt, vup vec3, vfov, aspectRatio, aperture, focusDist float64) *camera {
	theta := degreesToRadians(vfov)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := unitVector(subtractVec3(lookFrom, lookAt))
	u := unitVector(cross(vup, w))
	v := cross(w, u)

	horizontal := multiplyVec3ByFactor(multiplyVec3ByFactor(u, viewportWidth), focusDist)
	vertical := multiplyVec3ByFactor(multiplyVec3ByFactor(v, viewportHeight), focusDist)
	lowerLeftCorner := subtractVec3(subtractVec3(subtractVec3(lookFrom, divideVec3(horizontal, 2)), divideVec3(vertical, 2.0)), multiplyVec3ByFactor(w, focusDist))
	lensRadius := aperture / 2

	return &camera{lookFrom, lowerLeftCorner, horizontal, vertical, u, v, w, lensRadius}
}

func (c *camera) getRay(s, t float64) ray {
	rd := multiplyVec3ByFactor(randomInUnitDisk(), c.lensRadius)
	offset := addVec3(multiplyVec3ByFactor(c.u, rd.x()), multiplyVec3ByFactor(c.v, rd.y()))

	return newRay(addVec3(c.origin, offset), subtractVec3(subtractVec3(addVec3(addVec3(c.lowerLeftCorner, multiplyVec3ByFactor(c.horizontal, s)), multiplyVec3ByFactor(c.vertical, t)), c.origin), offset))
}
