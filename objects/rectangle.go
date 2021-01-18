package objects

import (
	"github.com/vfrazao-ns1/raytracing1weekend/ray"
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

// Rectangle is a 2d rectangle in 3d space
type Rectangle struct {
	A   vec3.Point // A first corner of rectangle
	W   vec3.Point // W second corner of rectangle
	H   vec3.Vec3  // H vector giving "height" to rectangle, simply summed to points A and W to get our remaining corners
	t1  Triangle   // Internal triangle representing half of the rectangle
	t2  Triangle   // Internal triangle representing half of the rectangle
	Mat Material   // Mat material is the triangle is made of
}

// InitRectangle initializes the internal triangles used to represent the rectangle
func (r *Rectangle) InitRectangle() {
	/*
		We represent rectangles as two triangles that share two of its vertices.
		Vector H should be at 90 degrees to AW but nothing enforces that so you
		could technically make some parallelograms.

		Illustration in 2d coordinates
		A = (2,2,0)
		W = (10,2,0)
		H = (0,5,0)

			A+H		         W+H
		7|  *----------------*
		6|  |    +           |
		5|  |      +         |
		4|  |        +       |
		3|  |          +     |
		2|  *----------------*
		1|  A                W
		 |____________________
		  1 2 3 4 5 6 7 8 9 10
	*/
	r.t1 = Triangle{
		V0:  r.A,
		V1:  r.W,
		V2:  r.A.Add(r.H),
		Mat: r.Mat,
	}
	r.t1.ComputeEdgesNormal()

	r.t2 = Triangle{
		V0:  r.t1.V2,
		V1:  r.t1.V1,
		V2:  r.t1.V1.Add(r.H),
		Mat: r.Mat,
	}
	r.t2.ComputeEdgesNormal()
}

func newRectangle(obj map[string]interface{}) (*Rectangle, error) {
	r := Rectangle{}
	var err error

	if c, ok := obj["a"].(map[string]interface{}); ok {
		r.A = vec3.Point{
			X: c["x"].(float64),
			Y: c["y"].(float64),
			Z: c["z"].(float64),
		}
	}
	if c, ok := obj["w"].(map[string]interface{}); ok {
		r.W = vec3.Point{
			X: c["x"].(float64),
			Y: c["y"].(float64),
			Z: c["z"].(float64),
		}
	}
	if c, ok := obj["h"].(map[string]interface{}); ok {
		r.H = vec3.Vec3{
			X: c["x"].(float64),
			Y: c["y"].(float64),
			Z: c["z"].(float64),
		}
	}

	if matInter, ok := obj["mat"].(map[string]interface{}); ok {
		r.Mat, err = newMaterial(matInter)
		if err != nil {
			return nil, err
		}
	}

	r.InitRectangle()

	return &r, nil
}

// Hit checks if a ray intersects with the triangle
func (r Rectangle) Hit(ray ray.Ray, tmin float64, tmax float64, rec *HitRecord) bool {
	if hit := r.t1.Hit(ray, tmin, tmax, rec); hit {
		return true
	}
	if hit := r.t2.Hit(ray, tmin, tmax, rec); hit {
		return true
	}
	return false
}
