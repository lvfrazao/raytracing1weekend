package objects

import (
	"github.com/vfrazao-ns1/raytracing1weekend/ray"
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

const (
	tmin = -1
	tmax = 1
)

// Hittable describes objects that can intersect rays
type Hittable interface {
	Hit(ray.Ray) float64
}

// HitRecord stores information on
type HitRecord struct {
	P      vec3.Point
	Normal vec3.Vec3
	T      float64
}
