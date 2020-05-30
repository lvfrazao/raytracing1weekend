package camera

import (
	"github.com/vfrazao-ns1/raytracing1weekend/ray"
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

const (
	aspect         = 16.0 / 9.0
	viewportHeight = 2.0
	focalLength    = 1.0
)

// Camera struct to represent our camera
type Camera struct {
	AspectRatio    float64 // AspectRatio usually 16:9
	ViewPortHeight float64 // ViewPortHeight usually 2.0
	ViewPortWidth  float64 // ViewPortWidth constrained by the aspect ratio and height
	FocalLength    float64 // FocalLength usually 1.0

	Origin     vec3.Point // Origin usually (0, 0, 0)
	Horizontal vec3.Vec3  // Horizontal line: <width, 0, 0>
	Vertical   vec3.Vec3  // Vertical line <0, height, 0>
	LowerLeft  vec3.Point // LowerLeft corner of view
}

// InitCamera Creates and initializes a Camera struct
func InitCamera() *Camera {
	c := new(Camera)

	c.AspectRatio = aspect
	c.ViewPortHeight = viewportHeight
	c.ViewPortWidth = c.AspectRatio * c.ViewPortHeight
	c.FocalLength = focalLength

	c.Origin = vec3.Point{X: 0, Y: 0, Z: 0}
	c.Horizontal = vec3.Vec3{X: c.ViewPortWidth, Y: 0, Z: 0}
	c.Vertical = vec3.Vec3{X: 0, Y: c.ViewPortHeight, Z: 0}
	c.LowerLeft = c.Origin.Sub(c.Horizontal.ScalarDiv(2)).Sub(c.Vertical.ScalarDiv(2)).Sub(vec3.Vec3{X: 0, Y: 0, Z: c.FocalLength})

	return c
}

// GetRay returns the ray from our camera
func (c Camera) GetRay(u, v float64) ray.Ray {
	return ray.Ray{Origin: c.Origin, Direction: c.LowerLeft.Add((c.Horizontal.ScalarMul(u))).Add(c.Vertical.ScalarMul(v).Sub(c.Origin))}
}
