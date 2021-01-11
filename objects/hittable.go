package objects

import (
	"encoding/json"

	"github.com/vfrazao-ns1/raytracing1weekend/ray"
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
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
	Material  Material   // Material that the object is made from
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
	hitAnything := false
	closest := tmax

	for _, obj := range hl.Data {
		if obj.Hit(r, tmin, closest, rec) {
			hitAnything = true
			closest = rec.T
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

// Hittables is a slice of hittables
type Hittables []Hittable

// UnmarshalJSON is just a custom unmarshaler for Hittables
// From: https://stackoverflow.com/questions/42721732/is-there-a-way-to-have-json-unmarshal-select-struct-type-based-on-type-prope
func (hs *Hittables) UnmarshalJSON(data []byte) error {
	// this just splits up the JSON array into the raw JSON for each object
	var raw []json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	for _, r := range raw {
		// unamrshal into a map to check the "type" field
		var obj map[string]interface{}
		err := json.Unmarshal(r, &obj)
		if err != nil {
			return err
		}

		hittableType := ""
		if t, ok := obj["type"].(string); ok {
			hittableType = t
		}

		// Send to custom constructor functions to instantiate object
		var actual Hittable
		switch hittableType {
		case "sphere":
			actual, err = newSphere(obj)
			if err != nil {
				return err
			}
		}

		*hs = append(*hs, actual)

	}
	return nil
}
