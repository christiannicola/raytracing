package main

import (
	"math"
	"math/rand"
)

var infinity = math.Inf(1)

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func randomFloat64() float64 {
	return rand.Float64()
}

func randomFloat64MinMax(min, max float64) float64 {
	return min + (max-min)*randomFloat64()
}
