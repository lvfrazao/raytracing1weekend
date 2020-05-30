package main

import (
	"math"
	"math/rand"

	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
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

func randomVec3() vec3.Vec3 {
	return vec3.Vec3{X: randomDouble(), Y: randomDouble(), Z: randomDouble()}
}

func randomVec3Between(min, max float64) vec3.Vec3 {
	return vec3.Vec3{X: randomDoubleBetween(min, max), Y: randomDoubleBetween(min, max), Z: randomDoubleBetween(min, max)}
}

func randomVec3InUnitSphere() vec3.Vec3 {
	for {
		p := randomVec3Between(-1, 1)
		if p.LengthSquared() >= 1 {
			return p
		}
	}
}

func randomUnitVector() vec3.Vec3 {
	a := randomDoubleBetween(0, 2*math.Pi)
	z := randomDoubleBetween(-1, 1)
	r := math.Sqrt(1 - z*z)
	return vec3.Vec3{X: r * math.Cos(a), Y: r * math.Sin(a), Z: z}
}

func randomInHemisphere(normal vec3.Vec3) vec3.Vec3 {
	inUnitSphere := randomVec3InUnitSphere()

	if inUnitSphere.Dot(normal) > 0 {
		return inUnitSphere
	}
	return inUnitSphere.Negate()
}
