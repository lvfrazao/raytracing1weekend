package vec3

import "fmt"

// Color represents a color
type Color = Vec3

// ColorString implements stringer for Color type
func (c Color) ColorString(samplesPerPixel int) string {
	r := c.X
	g := c.Y
	b := c.Z

	scale := 1.0 / float64(samplesPerPixel)
	r *= scale
	g *= scale
	b *= scale

	return fmt.Sprintf("%d %d %d", int(256*clamp(r, 0, 0.999)), int(256*clamp(g, 0, 0.999)), int(256*clamp(b, 0, 0.999)))
}
