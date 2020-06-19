package renderer

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/vfrazao-ns1/raytracing1weekend/utils"
)

type PNGRenderer struct {
	ImageWidth      int
	ImageHeight     int
	ImagePixels     []Pixel
	SamplesPerPixel int
}

func (p PNGRenderer) Render(fname string) {
	renderFile, err := os.Create(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create file, %s, for rendering: %v\n", fname, err)
		os.Exit(1)
	}
	defer renderFile.Close()

	img := image.NewNRGBA(image.Rect(0, 0, p.ImageWidth, p.ImageHeight))

	for _, pixel := range p.ImagePixels {
		r, g, b, _ := pixel.Color.RGBA(p.SamplesPerPixel)
		img.Set(int(pixel.Position.X), int(pixel.Position.Y), color.NRGBA{
			R: uint8(255 * utils.Clamp(r, 0, 0.999)),
			G: uint8(255 * utils.Clamp(g, 0, 0.999)),
			B: uint8(255 * utils.Clamp(b, 0, 0.999)),
			A: 255,
		})
	}

	png.Encode(renderFile, img)
}
