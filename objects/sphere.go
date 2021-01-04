package objects

import (
	"math"

	"github.com/vfrazao-ns1/raytracing1weekend/ray"
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

// Sphere represents a sphere in space
type Sphere struct {
	Center vec3.Point // Center of sphere
	Radius float64    // Radius of sphere
	Mat    Material   // Mat material the sphere is made of
}

func newSphere(obj map[string]interface{}) (*Sphere, error) {
	// I hate this so so much
	s := Sphere{}
	var err error

	if c, ok := obj["center"].(map[string]interface{}); ok {
		s.Center = vec3.Point{
			X: c["x"].(float64),
			Y: c["y"].(float64),
			Z: c["z"].(float64),
		}
	}

	if r, ok := obj["radius"].(float64); ok {
		s.Radius = r
	}

	if matInter, ok := obj["mat"].(map[string]interface{}); ok {
		s.Mat, err = newMaterial(matInter)
		if err != nil {
			return nil, err
		}
	}

	return &s, nil
}

// Hit checks if a ray intersects with the sphere
func (s Sphere) Hit(ray ray.Ray, tmin float64, tmax float64, rec *HitRecord) bool {
	// vector origin -> center
	oc := ray.Origin.Sub(s.Center)
	a := ray.Direction.LengthSquared()
	halfB := oc.Dot(ray.Direction)
	c := oc.LengthSquared() - s.Radius*s.Radius
	discriminant := halfB*halfB - a*c
	if discriminant > 0 {
		root := math.Sqrt(discriminant)
		temp := (-halfB - root) / a
		if temp < tmax && temp > tmin {
			rec.T = temp
			rec.P = ray.Position(rec.T)
			outwardNormal := rec.P.Sub(s.Center).ScalarDiv(s.Radius)
			rec.SetFaceNormal(ray, outwardNormal)
			rec.Material = s.Mat
			return true
		}
		temp = (-halfB + root) / a
		if temp < tmax && temp > tmin {
			rec.T = temp
			rec.P = ray.Position(rec.T)
			outwardNormal := rec.P.Sub(s.Center).ScalarDiv(s.Radius)
			rec.SetFaceNormal(ray, outwardNormal)
			rec.Material = s.Mat
			return true
		}
	}
	return false
}
