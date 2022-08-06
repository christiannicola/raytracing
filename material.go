package main

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
)

func newLambertian(color vec3) lambertian {
	return lambertian{color}
}

func (l lambertian) scatter(rayIn *ray, rec *hitRecord, attenuation *vec3, scattered *ray) bool {
	scatterDirection := addVec3(rec.normal, randomUnitVector())

	if scatterDirection.neadZero() {
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
