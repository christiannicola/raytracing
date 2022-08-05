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
