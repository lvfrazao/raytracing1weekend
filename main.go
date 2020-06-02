package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/vfrazao-ns1/raytracing1weekend/camera"
	"github.com/vfrazao-ns1/raytracing1weekend/objects"
	"github.com/vfrazao-ns1/raytracing1weekend/ray"
	"github.com/vfrazao-ns1/raytracing1weekend/utils"
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

const (
	maxColor = 255
	fileName = "render1.ppm"
)

func main() {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error openning file for render: %s", err)
		os.Exit(1)
	}
	defer f.Close()

	imgWidth := 384
	aspect := 16.0 / 9.0
	imgHeight := int(float64(imgWidth) / aspect)
	samplesPerPixel := 100
	maxDepth := 50

	numPixels := (imgHeight * imgWidth)
	pixels := make([]string, numPixels)

	lookfrom := vec3.Point{X: 3, Y: 3, Z: 2}
	lookat := vec3.Point{X: 0, Y: 0, Z: -1}
	vup := vec3.Vec3{X: 0, Y: 1, Z: 0}
	distToFocus := lookfrom.Sub(lookat).Length()
	aperture := 2.0
	cam := camera.InitCamera(lookfrom, lookat, vup, 20, float64(imgWidth)/float64(imgHeight), aperture, distToFocus)
	world := new(objects.HittableList)
	world.Add(
		objects.Sphere{
			Center: vec3.Point{X: 0, Y: 0, Z: -1},
			Radius: 0.5,
			Mat: objects.Lambertian{
				Albedo: vec3.Color{X: 0.1, Y: 0.2, Z: 0.5},
			},
			// Mat: objects.DiElectric{
			// 	RefIndex: 1.5, // Glass
			// },
		},
	)

	world.Add(
		objects.Sphere{
			Center: vec3.Point{X: 0, Y: -100.5, Z: -1},
			Radius: 100,
			Mat: objects.Lambertian{
				Albedo: vec3.Color{X: 0.8, Y: 0.8, Z: 0},
			},
		},
	)

	world.Add(
		objects.Sphere{
			Center: vec3.Point{X: 1, Y: 0, Z: -1},
			Radius: 0.5,
			Mat: objects.Metal{
				Albedo: vec3.Color{X: 0.8, Y: 0.6, Z: 0.2},
				Fuzz:   0.0,
			},
			// Mat: objects.DiElectric{
			// 	RefIndex: 1.5, // Glass
			// },
		},
	)

	world.Add(
		objects.Sphere{
			Center: vec3.Point{X: -1, Y: 0, Z: -1},
			Radius: 0.5,
			// Mat: objects.Metal{
			// 	Albedo: vec3.Color{X: 0.9, Y: 0.9, Z: 0.9},
			// 	Fuzz:   0.2,
			// },
			Mat: objects.DiElectric{
				RefIndex: 1.5, // Glass
			},
		},
	)

	world.Add(
		objects.Sphere{
			Center: vec3.Point{X: -1, Y: 0, Z: -1},
			Radius: -0.45,
			Mat: objects.DiElectric{
				RefIndex: 1.5, // Glass
			},
		},
	)

	fmt.Fprintf(f, "P3\n")
	fmt.Fprintf(f, "%d %d\n", imgWidth, imgHeight)
	fmt.Fprintf(f, "%d\n", maxColor)
	for j := (imgHeight - 1); j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %10d", j)
		for i := 0; i < imgWidth; i++ {

			pixel := vec3.Color{X: 0, Y: 0, Z: 0}
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + utils.RandomDouble()) / float64(imgWidth-1)
				v := (float64(j) + utils.RandomDouble()) / float64(imgHeight-1)
				ray := cam.GetRay(u, v)
				pixel = pixel.Add(RayColor(ray, world, maxDepth))
			}

			pixels[(imgWidth*(imgHeight-1-j))+i] = pixel.ColorString(samplesPerPixel)
		}
	}
	fmt.Fprintf(f, strings.Join(pixels, "\n"))
	fmt.Fprintf(os.Stderr, "\nDone\n")
}

// RayColor returns the ray color
func RayColor(r ray.Ray, world objects.Hittable, depth int) vec3.Color {
	hitRec := new(objects.HitRecord)

	if depth <= 0 {
		return vec3.Color{X: 0, Y: 0, Z: 0}
	}

	tmin := 0.001
	tmax := math.Inf(1)
	if world.Hit(r, tmin, tmax, hitRec) {
		scattered := new(ray.Ray)
		attenuation := new(vec3.Color)
		if hitRec.Material.Scatter(r, *hitRec, attenuation, scattered) {
			return attenuation.Mul(RayColor(*scattered, world, depth-1))
		}
		return vec3.Color{X: 0, Y: 0, Z: 0}
	}

	// If no hits then the color == background
	unitDirection := r.Direction.Unit()
	t := 0.5 * (unitDirection.Y + 1.0)
	return vec3.Color{X: 1, Y: 1, Z: 1}.ScalarMul(1 - t).Add(vec3.Color{X: 0.5, Y: 0.7, Z: 1}.ScalarMul(t))
}
