package main

import (
	"fmt"
	"io"
	"math"
)

func writeColor(w io.Writer, pixelColor vec3, samplesPerPixel int) error {
	scale := 1.0 / float64(samplesPerPixel)

	r := clamp(math.Sqrt(pixelColor.x()*scale), 0.0, 0.999) * 256
	g := clamp(math.Sqrt(pixelColor.y()*scale), 0.0, 0.999) * 256
	b := clamp(math.Sqrt(pixelColor.z()*scale), 0.0, 0.999) * 256

	_, err := fmt.Fprintf(w, "%d %d %d\n", int(r), int(g), int(b))

	return err
}

func rayColor(r *ray, world hittable, depth int) vec3 {
	rec := newHitRecord()

	if depth <= 0 {
		return emptyVec3()
	}

	if world.hit(r, 0.001, infinity, &rec) {
		scattered := ray{}
		attenuation := emptyVec3()

		if rec.material.scatter(r, &rec, &attenuation, &scattered) {
			return multiplyVec3(attenuation, rayColor(&scattered, world, depth-1))
		}

		return emptyVec3()
	}

	unitDirection := unitVector(r.direction)
	t := 0.5 * (unitDirection.y() + 1.0)

	return addVec3(multiplyVec3ByFactor(newVec3(1.0, 1.0, 1.0), 1.0-t), multiplyVec3ByFactor(newVec3(0.5, 0.7, 1.0), t))
}

func clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}

	if x > max {
		return max
	}

	return x
}
