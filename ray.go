package main

type ray struct {
	origin    vec3
	direction vec3
}

func newRay(origin, direction vec3) ray {
	return ray{origin, direction}
}

func (r ray) at(t float64) vec3 {
	r.direction.multiply(t)
	r.origin = addVec3(r.origin, r.direction)

	return r.origin
}
