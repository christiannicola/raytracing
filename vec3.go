package main

import (
	"fmt"
	"io"
	math "math"
)

type vec3 [3]float64

func newVec3(e0, e1, e2 float64) vec3 {
	return [3]float64{e0, e1, e2}
}

func emptyVec3() vec3 {
	return [3]float64{}
}

func (v *vec3) x() float64 {
	return v[0]
}

func (v *vec3) y() float64 {
	return v[1]
}

func (v *vec3) z() float64 {
	return v[2]
}

func (v *vec3) negate() vec3 {
	return newVec3(-v.x(), -v.y(), -v.z())
}

func (v *vec3) add(f float64) {
	v[0] += f
	v[1] += f
	v[2] += f
}

func (v *vec3) multiply(f float64) {
	v[0] *= f
	v[1] *= f
	v[2] *= f
}

func (v *vec3) divide(d float64) {
	v.multiply(1 / d)
}

func (v vec3) length() float64 {
	return math.Sqrt(v.lengthSquared())
}

func (v *vec3) lengthSquared() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

func (v *vec3) nearZero() bool {
	s := 1e-8

	return (math.Abs(v.x()) < s) && (math.Abs(v.y()) < s) && (math.Abs(v.z()) < s)
}

func addVec3(lhs, rhs vec3) vec3 {
	return newVec3(lhs.x()+rhs.x(), lhs.y()+rhs.y(), lhs.z()+rhs.z())
}

func subtractVec3(lhs, rhs vec3) vec3 {
	return newVec3(lhs.x()-rhs.x(), lhs.y()-rhs.y(), lhs.z()-rhs.z())
}

func multiplyVec3(lhs, rhs vec3) vec3 {
	return newVec3(lhs.x()*rhs.x(), lhs.y()*rhs.y(), lhs.z()*rhs.z())
}

func multiplyVec3ByFactor(v vec3, f float64) vec3 {
	v.multiply(f)

	return v
}

func divideVec3(v vec3, d float64) vec3 {
	v.divide(d)

	return v
}

func (v *vec3) debugPrint(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%f %f %f", v.x(), v.y(), v.z())

	return err
}

func dot(u vec3, v vec3) float64 {
	return u.x()*v.x() + u.y()*v.y() + u.z()*v.z()
}

func cross(u vec3, v vec3) vec3 {
	return newVec3(u.y()*v.z()-u.z()*v.y(), u.z()*v.x()-u.x()*v.z(), u.x()*v.y()-u.y()*v.x())
}

func unitVector(v vec3) vec3 {
	v.divide(v.length())

	return v
}

func random() vec3 {
	return newVec3(randomFloat64(), randomFloat64(), randomFloat64())
}

func randomMinMax(min, max float64) vec3 {
	return newVec3(randomFloat64MinMax(min, max), randomFloat64MinMax(min, max), randomFloat64MinMax(min, max))
}

func randomInUnitSphere() vec3 {
	for {
		p := random()
		if p.lengthSquared() < 1 {
			return p
		}
	}
}

func randomUnitVector() vec3 {
	return unitVector(randomInUnitSphere())
}

func randomInHemisphere(normal vec3) vec3 {
	inUnitSphere := randomInUnitSphere()

	if dot(inUnitSphere, normal) > 0.0 {
		return inUnitSphere
	}

	return inUnitSphere.negate()
}

func reflect(v, n vec3) vec3 {
	return subtractVec3(v, multiplyVec3ByFactor(n, 2*dot(v, n)))
}

func refract(uv, n vec3, etaiOverEtat float64) vec3 {
	cosTheta := math.Min(dot(uv.negate(), n), 1.0)

	rOutPerp := multiplyVec3ByFactor(addVec3(multiplyVec3ByFactor(n, cosTheta), uv), etaiOverEtat)
	rOutParallel := multiplyVec3ByFactor(n, -(math.Sqrt(math.Abs(1.0 - rOutPerp.lengthSquared()))))

	return addVec3(rOutParallel, rOutPerp)
}

func randomInUnitDisk() vec3 {
	for {
		p := newVec3(randomFloat64MinMax(-1, 1), randomFloat64MinMax(-1, 1), 0)

		if p.lengthSquared() >= 1 {
			continue
		}

		return p
	}
}
