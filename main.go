package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/vfrazao-ns1/raytracing1weekend/objects"
	"github.com/vfrazao-ns1/raytracing1weekend/ray"
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

const (
	maxColor = 255
	fileName = "render.ppm"
)

func main() {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error openning file for render: %s", err)
		os.Exit(1)
	}
	defer f.Close()

	imgWidth := 3840
	aspect := 16.0 / 9.0
	imgHeight := int(float64(imgWidth) / aspect)
	numPixels := (imgHeight * imgWidth)
	pixels := make([]string, numPixels)

	origin := vec3.Point{0, 0, 0}
	horizontal := vec3.Point{4, 0, 0}
	vertical := vec3.Point{0, 2.25, 0}

	lowerLeftCorner := origin.Sub(horizontal.ScalarDiv(2)).Sub(vertical.ScalarDiv(2)).Sub(vec3.Vec3{0, 0, 1})

	fmt.Fprintf(f, "P3\n")
	fmt.Fprintf(f, "%d %d\n", imgWidth, imgHeight)
	fmt.Fprintf(f, "%d\n", maxColor)
	for j := (imgHeight - 1); j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %10d", j)
		for i := 0; i < imgWidth; i++ {
			u := float64(i) / float64(imgWidth-1)
			v := float64(j) / float64(imgHeight-1)
			ray := ray.Ray{
				Origin:    origin,
				Direction: lowerLeftCorner.Add(horizontal.ScalarMul(u)).Add(vertical.ScalarMul(v)),
			}

			pixels[(imgWidth*(imgHeight-1-j))+i] = RayColor(ray).ColorString()
		}
	}
	fmt.Fprintf(f, strings.Join(pixels, "\n"))
	fmt.Fprintf(os.Stderr, "\nDone\n")
}

// RayColor returns the ray color
func RayColor(r ray.Ray) vec3.Color {
	sphere := objects.Sphere{Center: vec3.Point{0, 0, -1}, Radius: 0.5}
	hitRec := objects.HitRecord{}
	_ = sphere.Hit(r, &hitRec)
	if hitRec.T > 0 {
		N := hitRec.P.Sub(vec3.Vec3{0, 0, -1})
		return vec3.Color{
			X: N.X,
			Y: N.Y,
			Z: N.Z,
		}.ScalarAdd(1).ScalarDiv(2)
	}
	direction := r.Direction.Unit()
	t := 0.5 * (direction.Y + 1.0)
	rayColor := vec3.Color{1, 1, 1}.ScalarMul(1 - t).Add(vec3.Color{0.5, 0.7, 1}.ScalarMul(t))
	return rayColor
}
