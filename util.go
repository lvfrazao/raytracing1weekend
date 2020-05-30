package main

import (
	"math"
	"math/rand"
)

func degrees2radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func randomDouble() float64 {
	return rand.Float64()
}

func randomDoubleBetween(min, max float64) float64 {
	return min + (max-min)*randomDouble()
}

func clamp(x, min, max float64) float64 {
	// Clamps the x value between min and max
	if x < min {
		return min
	} else if x > max {
		return max
	}
	return x
}
