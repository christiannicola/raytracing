package main

import (
	"fmt"
	"io"
	"math"
)

func writeColor(w io.Writer, pixelColor vec3) error {
	r := int(255.999 * pixelColor.x())
	g := int(255.999 * pixelColor.y())
	b := int(255.999 * pixelColor.z())

	_, err := fmt.Fprintf(w, "%d %d %d\n", r, g, b)

	return err
}

func rayColor(r *ray) vec3 {
	var target float64

	if target = hitSphere(newVec3(0, 0, -1), 0.5, r); target > 0 {
		n := unitVector(subtractVec3(r.at(target), newVec3(0, 0, -1)))

		return multiplyVec3ByFactor(newVec3(n.x()+1, n.y()+1, n.z()+1), 0.5)
	}

	unitDirection := unitVector(r.direction)

	target = 0.5 * (unitDirection.y() + 1.0)

	whiteColor := multiplyVec3ByFactor(newVec3(1.0, 1.0, 1.0), 1.0-target)
	blueColor := multiplyVec3ByFactor(newVec3(0.5, 0.7, 1.0), target)

	return addVec3(whiteColor, blueColor)
}

func hitSphere(center vec3, radius float64, r *ray) float64 {
	originCenter := subtractVec3(r.origin, center)

	a := r.direction.lengthSquared()
	halfB := dot(originCenter, r.direction)
	c := originCenter.lengthSquared() - radius*radius

	discriminant := halfB*halfB - a*c

	if discriminant < 0 {
		return -1.0
	}

	return (-halfB - math.Sqrt(discriminant)) / a
}
