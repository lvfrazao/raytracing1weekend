package vec3

import (
	"fmt"
	"math"
)

// Color represents a color
type Color = Vec3

// ColorString implements stringer for Color type
func (c Color) ColorString(samplesPerPixel int) string {
	r, g, b, _ := c.RGBA(samplesPerPixel)
	return fmt.Sprintf("%d %d %d", int(256*clamp(r, 0, 0.999)), int(256*clamp(g, 0, 0.999)), int(256*clamp(b, 0, 0.999)))
}

func (c Color) RGBA(samplesPerPixel int) (float64, float64, float64, float64) {
	r := c.X
	g := c.Y
	b := c.Z

	scale := 1.0 / float64(samplesPerPixel)
	r = math.Sqrt(scale * r)
	g = math.Sqrt(scale * g)
	b = math.Sqrt(scale * b)
	return r, g, b, 1
}
