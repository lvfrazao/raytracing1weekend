package renderer

import (
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

type Pixel struct {
	Color    vec3.Color
	Position vec3.Point
}

type Renderer interface {
	Render([]Pixel)
}
