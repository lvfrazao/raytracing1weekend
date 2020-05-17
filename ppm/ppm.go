package ppm

import (
	"fmt"

	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

// PPM image format
type PPM struct {
	Width    int
	Height   int
	MaxColor int
	Pixels   []vec3.Color
}

// PrintFile prints to stdout
func (p PPM) PrintFile() {
	fmt.Printf("P3\n")
	fmt.Printf("%d %d\n", p.Width, p.Height)
	fmt.Printf("%d\n", p.MaxColor)
	for _, pixel := range p.Pixels {
		fmt.Printf("%s\n", pixel.ColorString())
	}
}
