package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/vfrazao-ns1/raytracing1weekend/objects"
	"github.com/vfrazao-ns1/raytracing1weekend/utils"
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

func staticScene() objects.HittableList {
	world := new(objects.HittableList)
	world.Add(
		objects.Sphere{
			Center: vec3.Point{X: 0, Y: 0, Z: -1},
			Radius: 0.5,
			Mat: objects.Lambertian{
				Albedo: vec3.Color{X: 0.1, Y: 0.2, Z: 0.5},
			},
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
		},
	)

	world.Add(
		objects.Sphere{
			Center: vec3.Point{X: -1, Y: 0, Z: -1},
			Radius: 0.5,
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
	return *world
}

func RandomWorld() objects.HittableList {
	randSeed := time.Now().UnixNano()
	fmt.Fprintf(os.Stderr, "Using %v as random seed\n", randSeed)
	rand.Seed(randSeed)

	world := new(objects.HittableList)

	ground := objects.Sphere{
		Center: vec3.Point{X: 0, Y: -1000, Z: 0},
		Radius: 1000,
		Mat:    objects.Lambertian{Albedo: vec3.Color{X: 0.5, Y: 0.5, Z: 0.5}},
	}

	world.Add(ground)

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := utils.RandomDouble()

			center := vec3.Point{
				X: float64(a) + 0.9*utils.RandomDouble(),
				Y: 0.2,
				Z: float64(b) + 0.9*utils.RandomDouble(),
			}

			if center.Sub(vec3.Vec3{X: 4, Y: 0.2, Z: 0}).Length() > 0.9 {
				var sphereMaterial objects.Material

				switch {
				case chooseMat < 0.8:
					// diffuse
					albedo := utils.RandomVec3().Mul(utils.RandomVec3())
					sphereMaterial = objects.Lambertian{Albedo: albedo}
				case chooseMat < 0.95:
					albedo := utils.RandomVec3Between(0.5, 1)
					fuzz := utils.RandomDoubleBetween(0, 0.5)
					sphereMaterial = objects.Metal{Albedo: albedo, Fuzz: fuzz}
				default:
					sphereMaterial = objects.DiElectric{1.5}

				}
				world.Add(
					objects.Sphere{
						Center: center,
						Radius: 0.2,
						Mat:    sphereMaterial,
					},
				)
			}
		}
	}

	world.Add(
		objects.Sphere{
			Center: vec3.Point{X: 0, Y: 1, Z: 0},
			Radius: 1.0,
			Mat:    objects.DiElectric{RefIndex: 1.5},
		},
	)

	world.Add(
		objects.Sphere{
			Center: vec3.Point{X: -4, Y: 1, Z: 0},
			Radius: 1.0,
			Mat:    objects.Lambertian{Albedo: vec3.Color{X: 0.4, Y: 0.2, Z: 0.1}},
		},
	)

	world.Add(
		objects.Sphere{
			Center: vec3.Point{X: 4, Y: 1, Z: 0},
			Radius: 1.0,
			Mat: objects.Metal{
				Albedo: vec3.Color{X: 0.7, Y: 0.6, Z: 0.5},
				Fuzz:   0.0,
			},
		},
	)

	return *world
}
