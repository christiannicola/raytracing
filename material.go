package main

import "math"

type (
	material interface {
		scatter(rayIn *ray, rec *hitRecord, attenuation *vec3, scattered *ray) bool
	}

	lambertian struct {
		albedo vec3
	}

	metal struct {
		albedo vec3
		fuzz   float64
	}

	dielectric struct {
		ir float64
	}
)

func newLambertian(color vec3) lambertian {
	return lambertian{color}
}

func (l lambertian) scatter(rayIn *ray, rec *hitRecord, attenuation *vec3, scattered *ray) bool {
	scatterDirection := addVec3(rec.normal, randomUnitVector())

	if scatterDirection.nearZero() {
		scatterDirection = rec.normal
	}

	*scattered = newRay(rec.p, scatterDirection)
	*attenuation = l.albedo

	return true
}

func newMetal(color vec3, f float64) metal {
	return metal{color, f}
}

func (m metal) scatter(rayIn *ray, rec *hitRecord, attenuation *vec3, scattered *ray) bool {
	reflected := reflect(unitVector(rayIn.direction), rec.normal)
	*scattered = newRay(rec.p, addVec3(reflected, multiplyVec3ByFactor(randomInUnitSphere(), m.fuzz)))
	*attenuation = m.albedo

	return dot(scattered.direction, rec.normal) > 0
}

func newDielectric(ir float64) dielectric {
	return dielectric{ir}
}

func (d dielectric) scatter(rayIn *ray, rec *hitRecord, attenuation *vec3, scattered *ray) bool {
	var (
		refractionRatio float64
		direction       vec3
	)

	*attenuation = newVec3(1.0, 1.0, 1.0)

	if rec.frontFace {
		refractionRatio = 1.0 / d.ir
	} else {
		refractionRatio = d.ir
	}

	unitDirection := unitVector(rayIn.direction)
	cosTheta := math.Min(dot(unitDirection.negate(), rec.normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := refractionRatio*sinTheta > 1.0

	if cannotRefract || d.reflectance(cosTheta, refractionRatio) > randomFloat64() {
		direction = reflect(unitDirection, rec.normal)
	} else {
		direction = refract(unitDirection, rec.normal, refractionRatio)
	}

	*scattered = newRay(rec.p, direction)

	return true
}

func (d dielectric) reflectance(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0

	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}
