package utils

import (
	"math"
	"math/rand"

	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

func Degrees2radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func RandomDouble() float64 {
	return rand.Float64()
}

func RandomDoubleBetween(min, max float64) float64 {
	return min + (max-min)*RandomDouble()
}

func Clamp(x, min, max float64) float64 {
	// Clamps the x value between min and max
	if x < min {
		return min
	} else if x > max {
		return max
	}
	return x
}

func RandomVec3() vec3.Vec3 {
	return vec3.Vec3{X: RandomDouble(), Y: RandomDouble(), Z: RandomDouble()}
}

func RandomVec3Between(min, max float64) vec3.Vec3 {
	return vec3.Vec3{X: RandomDoubleBetween(min, max), Y: RandomDoubleBetween(min, max), Z: RandomDoubleBetween(min, max)}
}

func RandomVec3InUnitSphere() vec3.Vec3 {
	for {
		p := RandomVec3Between(-1, 1)
		if p.LengthSquared() >= 1 {
			return p
		}
	}
}

func RandomUnitVector() vec3.Vec3 {
	a := RandomDoubleBetween(0, 2*math.Pi)
	z := RandomDoubleBetween(-1, 1)
	r := math.Sqrt(1 - z*z)
	return vec3.Vec3{X: r * math.Cos(a), Y: r * math.Sin(a), Z: z}
}

func RandomInHemisphere(normal vec3.Vec3) vec3.Vec3 {
	inUnitSphere := RandomVec3InUnitSphere()

	if inUnitSphere.Dot(normal) > 0 {
		return inUnitSphere
	}
	return inUnitSphere.Negate()
}
