package objects

import (
	"math"

	"github.com/vfrazao-ns1/raytracing1weekend/ray"
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

// Triangle is a 2d triangle in 3d space
type Triangle struct {
	// V0, V1, V2 proceed in counterclockwise if you were looking straight on to the triangle
	V0     vec3.Point // Position of vertex A of a triangle
	V1     vec3.Point // Position of vertex A of a triangle
	V2     vec3.Point // Position of vertex A of a triangle
	A      vec3.Vec3  // v1 - v0 edge of triangle
	B      vec3.Vec3  // v2 - v0 edge of triangle
	C      vec3.Vec3  // v2 - v1 edge of triangle
	Normal vec3.Vec3  // v1 - v0 edge of triangle
	Mat    Material   // Mat material is the triangle is made of
}

func newTriangle(obj map[string]interface{}) (*Triangle, error) {
	t := Triangle{}
	var err error

	if c, ok := obj["v0"].(map[string]interface{}); ok {
		t.V0 = vec3.Point{
			X: c["x"].(float64),
			Y: c["y"].(float64),
			Z: c["z"].(float64),
		}
	}
	if c, ok := obj["v1"].(map[string]interface{}); ok {
		t.V1 = vec3.Point{
			X: c["x"].(float64),
			Y: c["y"].(float64),
			Z: c["z"].(float64),
		}
	}
	if c, ok := obj["v2"].(map[string]interface{}); ok {
		t.V2 = vec3.Point{
			X: c["x"].(float64),
			Y: c["y"].(float64),
			Z: c["z"].(float64),
		}
	}

	if matInter, ok := obj["mat"].(map[string]interface{}); ok {
		t.Mat, err = newMaterial(matInter)
		if err != nil {
			return nil, err
		}
	}

	t.ComputeEdgesNormal()

	return &t, nil
}

// ComputeEdgesNormal calculates the A, B, and C edges as well as the plane normal vector
func (t *Triangle) ComputeEdgesNormal() {
	t.A = t.V1.Sub(t.V0)
	t.B = t.V2.Sub(t.V0)
	t.C = t.V2.Sub(t.V1)
	t.Normal = t.A.Cross(t.B).Unit()
}

// Hit checks if a ray intersects with the triangle
func (t Triangle) Hit(ray ray.Ray, tmin float64, tmax float64, rec *HitRecord) bool {
	// Check if ray intersects triangle using the Möller-Trumbore algorithm
	// From here: https://www.scratchapixel.com/lessons/3d-basic-rendering/ray-tracing-rendering-a-triangle/moller-trumbore-ray-triangle-intersection
	pvec := ray.Direction.Cross(t.B)
	det := t.A.Dot(pvec)

	if math.Abs(det) < 0.00001 {
		//ray and triangle are parallel
		return false
	}
	invDet := 1 / det
	tvec := ray.Origin.Sub(t.V0)
	u := tvec.Dot(pvec) * invDet
	if u < 0 || u > 1 {
		return false
	}

	qvec := tvec.Cross(t.A)
	v := ray.Direction.Dot(qvec) * invDet
	if v < 0 || u+v > 1 {
		return false
	}

	// Time at which the ray intersect our plane
	tIntersect := t.B.Dot(qvec) * invDet

	// Check if the triangle is "behind" us
	if tIntersect < tmin || tIntersect > tmax {
		return false
	}

	// We now compute the point at which the ray intersects our plane
	pHit := ray.Position(tIntersect)
	rec.T = tIntersect
	rec.P = pHit
	rec.SetFaceNormal(ray, t.Normal)
	rec.Material = t.Mat
	return true
}
