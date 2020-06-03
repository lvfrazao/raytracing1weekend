package objects

import (
	"math"

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

// DiElectric materials like glass and water
type DiElectric struct {
	RefIndex float64
}

// Scatter implements `Material` interface for DiElectric
func (d DiElectric) Scatter(rIn ray.Ray, rec HitRecord, attenuation *vec3.Color, scattered *ray.Ray) bool {
	*attenuation = vec3.Color{X: 1, Y: 1, Z: 1}
	var etaiOverEtat float64

	if rec.FrontFace {
		etaiOverEtat = 1.0 / d.RefIndex
	} else {
		etaiOverEtat = d.RefIndex
	}
	unitDirection := rIn.Direction.Unit()

	cosTheta := utils.Fmin(unitDirection.Negate().Dot(rec.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	if etaiOverEtat*sinTheta > 1.0 || utils.RandomDouble() < d.schlick(cosTheta, etaiOverEtat) {
		reflected := vec3.Reflect(unitDirection, rec.Normal)
		scattered.Origin = rec.P
		scattered.Direction = reflected
		return true
	}

	refracted := vec3.Refract(unitDirection, rec.Normal, etaiOverEtat)
	scattered.Origin = rec.P
	scattered.Direction = refracted
	return true
}

func (d DiElectric) schlick(cosine, refindex float64) float64 {
	r0 := (1 - refindex) / (1 + refindex)
	r0 = r0 * r0
	return r0 * (1 - r0) * math.Pow((1-cosine), 5)
}
