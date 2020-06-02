package objects

import (
	"github.com/vfrazao-ns1/raytracing1weekend/ray"
	"github.com/vfrazao-ns1/raytracing1weekend/utils"
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

// A Material is an interface representing any possible material
type Material interface {
	Scatter(ray.Ray, HitRecord, *vec3.Color, *ray.Ray) bool
}

// Lambertian material type struct
type Lambertian struct {
	Albedo vec3.Color // Albedo of the material (basically how reflective it is)
}

// Scatter calculates the color attenuation and scattering
func (l Lambertian) Scatter(rIn ray.Ray, rec HitRecord, attenuation *vec3.Color, scattered *ray.Ray) bool {
	scatterDir := rec.Normal.Add(utils.RandomUnitVector())

	scattered.Origin = rec.P
	scattered.Direction = scatterDir

	*attenuation = l.Albedo
	return true
}

// Metal material type
type Metal struct {
	Albedo vec3.Color // Albedo of the material (basically how reflective it is)
	Fuzz   float64    // Fuzz iness of the reflections
}

// Scatter calculates the color attenuation and scattering
func (m Metal) Scatter(rIn ray.Ray, rec HitRecord, attenuation *vec3.Color, scattered *ray.Ray) bool {
	reflected := vec3.Reflect(rIn.Direction.Unit(), rec.Normal)

	scattered.Origin = rec.P
	scattered.Direction = reflected.Add(utils.RandomVec3InUnitSphere().ScalarMul(m.Fuzz))

	*attenuation = m.Albedo
	return scattered.Direction.Dot(rec.Normal) > 0
}
