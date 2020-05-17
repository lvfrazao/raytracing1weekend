package ray

import (
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

// Ray represents a ray, has an origin and direction
type Ray struct {
	Origin    vec3.Point
	Direction vec3.Vec3
}

// Position position of the ray at any given time t
func (r Ray) Position(t float64) vec3.Point {
	return r.Origin.Add(r.Direction.ScalarMul(t))
}
