package main

type camera struct {
	origin          vec3
	lowerLeftCorner vec3
	horizontal      vec3
	vertical        vec3
}

func newCamera() *camera {
	aspectRatio := 16.0 / 9.0
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := newVec3(0, 0, 0)
	horizontal := newVec3(viewportWidth, 0.0, 0.0)
	vertical := newVec3(0.0, viewportHeight, 0.0)
	lowerLeftCorner := subtractVec3(subtractVec3(subtractVec3(origin, divideVec3(horizontal, 2.0)), divideVec3(vertical, 2.0)), newVec3(0, 0, focalLength))

	return &camera{origin, lowerLeftCorner, horizontal, vertical}
}

func (c *camera) getRay(u, v float64) ray {
	direction := subtractVec3(addVec3(addVec3(c.lowerLeftCorner, multiplyVec3ByFactor(c.horizontal, u)), multiplyVec3ByFactor(c.vertical, v)), c.origin)

	return newRay(c.origin, direction)
}
