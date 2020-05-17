package vec3

import "fmt"

// Color represents a color
type Color = Vec3

// ColorString implements stringer for Color type
func (c Color) ColorString() string {
	return fmt.Sprintf("%d %d %d", int(c.X*255.999), int(c.Y*255.999), int(c.Z*255.999))
}
