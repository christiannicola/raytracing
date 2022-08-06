package main

import "math"

var infinity float64 = math.Inf(1)

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
