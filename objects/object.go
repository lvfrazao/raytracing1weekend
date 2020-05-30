package objects

import (
	"github.com/vfrazao-ns1/raytracing1weekend/ray"
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

const (
	tMin = -1
	tMax = 1
)

// Hittable describes objects that can intersect rays
type Hittable interface {
	Hit(ray.Ray, float64, float64, *HitRecord) bool
}

// HitRecord stores information on rays that have hit surfaces
type HitRecord struct {
	P         vec3.Point // P point where a surface was hit
	Normal    vec3.Vec3  // Normal vector to the hit surface
	T         float64    // T time at which the ray hit
	FrontFace bool       // FrontFace whether faces the front
}

// SetFaceNormal determines the normal to the surface
func (h *HitRecord) SetFaceNormal(r ray.Ray, outwardNormal vec3.Vec3) {
	h.FrontFace = r.Direction.Dot(outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Negate()
	}
}

// HittableList A list of hittable objects
type HittableList struct {
	Data []Hittable
}

// Hit were any hittables hit?
func (hl HittableList) Hit(r ray.Ray, tmin float64, tmax float64, rec *HitRecord) bool {
	tempRec := new(HitRecord)
	hitAnything := false
	closest := tmax

	for _, obj := range hl.Data {
		if obj.Hit(r, tmin, closest, tempRec) {
			hitAnything = true
			closest = tempRec.T
			*rec = *tempRec
		}
	}

	return hitAnything
}

// Add appends an Hittable element to the Data slice
func (hl *HittableList) Add(elem Hittable) {
	if hl.Data == nil {
		hl.Data = make([]Hittable, 0)
	}
	hl.Data = append(hl.Data, elem)
}
