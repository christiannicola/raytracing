package main

import (
	"fmt"
	"io"
	"math"
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

func (v *vec3) add(r vec3) {
	v[0] += r[0]
	v[1] += r[1]
	v[2] += r[2]
}

func (v *vec3) multiply(f float64) {
	v[0] *= f
	v[1] *= f
	v[2] *= f
}

func (v *vec3) divide(d float64) {
	v.multiply(1 / d)
}

func (v *vec3) length() float64 {
	return math.Sqrt(v.lengthSquared())
}

func (v *vec3) lengthSquared() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

func (v *vec3) debugPrint(w io.Writer) error {
	_, err := fmt.Fprintf(w, "%f %f %f", v.x(), v.y(), v.z())

	return err
}

func dot(u *vec3, v *vec3) float64 {
	return u.x()*v.x() + u.y()*v.y() + u.z()*v.z()
}

func cross(u *vec3, v *vec3) vec3 {
	return newVec3(u.y()*v.z()-u.z()*v.y(), u.z()*v.x()-u.x()*v.z(), u.x()*v.y()-u.y()*v.x())
}

func unitVector(v vec3) vec3 {
	v.divide(v.length())

	return v
}
