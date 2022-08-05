package main

import (
	"fmt"
	"io"
)

func writeColor(w io.Writer, pixelColor vec3) error {
	r := int(255.999 * pixelColor.x())
	g := int(255.999 * pixelColor.y())
	b := int(255.999 * pixelColor.z())

	_, err := fmt.Fprintf(w, "%d %d %d\n", r, g, b)

	return err
}

func rayColor(r *ray) vec3 {
	if hitSphere(newVec3(0, 0, -1), 0.5, r) {
		return newVec3(1, 0, 0)
	}

	unitDirection := unitVector(r.direction)

	t := 0.5 * (unitDirection.y() + 1.0)

	whiteColor := multiplyVec3ByFactor(newVec3(1.0, 1.0, 1.0), 1.0-t)
	blueColor := multiplyVec3ByFactor(newVec3(0.5, 0.7, 1.0), t)

	return addVec3(whiteColor, blueColor)
}

func hitSphere(center vec3, radius float64, r *ray) bool {
	originCenter := subtractVec3(r.origin, center)

	a := dot(r.direction, r.direction)
	b := dot(originCenter, r.direction) * 2.0
	c := dot(originCenter, originCenter) - radius*radius

	discriminant := b*b - 4*a*c

	return discriminant > 0
}
