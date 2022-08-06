package main

import (
	"math"
)

type (
	hittable interface {
		hit(r *ray, tMin, tMax float64, rec *hitRecord) bool
	}

	hitRecord struct {
		p         vec3
		normal    vec3
		t         float64
		frontFace bool
	}

	sphere struct {
		center vec3
		radius float64
	}

	hittableList struct {
		useConcurrency bool
		objects        []hittable
	}
)

func newHitRecord() hitRecord {
	return hitRecord{p: emptyVec3(), normal: emptyVec3()}
}

func (h *hitRecord) setFaceNormal(r *ray, outwardNormal vec3) {
	h.frontFace = dot(r.direction, outwardNormal) < 0

	if h.frontFace {
		h.normal = outwardNormal
	} else {
		h.normal = outwardNormal.negate()
	}
}

func newSphere(center vec3, radius float64) sphere {
	return sphere{center, radius}
}

func (s sphere) hit(r *ray, tMin, tMax float64, rec *hitRecord) bool {
	originCenter := subtractVec3(r.origin, s.center)

	a := r.direction.lengthSquared()
	halfB := dot(originCenter, r.direction)
	c := originCenter.lengthSquared() - s.radius*s.radius

	discriminant := halfB*halfB - a*c

	if discriminant < 0 {
		return false
	}

	sqrtD := math.Sqrt(discriminant)
	root := (-halfB - sqrtD) / a

	if root < tMin || tMax < root {
		root = (-halfB + sqrtD) / a
		if root < tMin || tMax < root {
			return false
		}
	}

	rec.t = root
	rec.p = r.at(rec.t)

	outwardNormal := divideVec3(subtractVec3(rec.p, s.center), s.radius)
	rec.setFaceNormal(r, outwardNormal)

	return true
}

func newHittableList() hittableList {
	l := hittableList{objects: []hittable{}}

	return l
}

func (l *hittableList) add(o hittable) {
	l.objects = append(l.objects, o)
}

func (l *hittableList) clear() {
	l.objects = []hittable{}
}

func (l *hittableList) hit(r *ray, tMin, tMax float64, rec *hitRecord) bool {
	tempRec := newHitRecord()
	hitAnything := false
	closestSoFar := tMax

	for i := range l.objects {
		if l.objects[i].hit(r, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.t
			*rec = tempRec
		}
	}

	return hitAnything
}
