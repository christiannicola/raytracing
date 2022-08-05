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
	unitDirection := unitVector(r.direction)

	t := 0.5 * (unitDirection.y() + 1.0)

	whiteColor := newVec3(1.0, 1.0, 1.0)
	whiteColor.multiply(1.0 - t)

	blueColor := newVec3(0.5, 0.7, 1.0)
	blueColor.multiply(t)

	whiteColor.add(blueColor)

	return whiteColor
}
