package main

import (
	"fmt"
	"math"
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

	origin := vec3.Point{X: 0, Y: 0, Z: 0}
	horizontal := vec3.Point{X: 4, Y: 0, Z: 0}
	vertical := vec3.Point{X: 0, Y: 2.25, Z: 0}

	lowerLeftCorner := origin.Sub(horizontal.ScalarDiv(2)).Sub(vertical.ScalarDiv(2)).Sub(vec3.Vec3{X: 0, Y: 0, Z: 1})

	world := new(objects.HittableList)
	world.Add(objects.Sphere{Center: vec3.Point{X: 0, Y: 0, Z: -1}, Radius: 0.5})
	world.Add(objects.Sphere{Center: vec3.Point{X: 0, Y: -100.5, Z: -1}, Radius: 100})

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

			pixels[(imgWidth*(imgHeight-1-j))+i] = RayColor(ray, world).ColorString()
		}
	}
	fmt.Fprintf(f, strings.Join(pixels, "\n"))
	fmt.Fprintf(os.Stderr, "\nDone\n")
}

// RayColor returns the ray color
func RayColor(r ray.Ray, world objects.Hittable) vec3.Color {
	hitRec := new(objects.HitRecord)
	tmin := 0.0
	tmax := math.Inf(1)
	if world.Hit(r, tmin, tmax, hitRec) {
		return vec3.Color{X: 1, Y: 1, Z: 1}.Add(hitRec.Normal).ScalarMul(0.5)
	}

	// If no hits then the color == background
	unitDirection := r.Direction.Unit()
	t := 0.5 * (unitDirection.Y + 1.0)
	return vec3.Color{X: 1, Y: 1, Z: 1}.ScalarMul(1 - t).Add(vec3.Color{X: 0.5, Y: 0.7, Z: 1}.ScalarMul(t))
}
