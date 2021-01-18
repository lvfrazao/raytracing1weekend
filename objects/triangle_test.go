package objects

import (
	"math"
	"testing"

	"github.com/vfrazao-ns1/raytracing1weekend/ray"
	"github.com/vfrazao-ns1/raytracing1weekend/vec3"
)

func isCloseEnough(a, b float64) bool {
	tolerance := 0.00001
	return math.Abs(a-b) < tolerance
}

func pointsEqual(p1, p2 vec3.Point) bool {
	if !isCloseEnough(p1.X, p2.X) {
		return false
	}
	if !isCloseEnough(p1.Y, p2.Y) {
		return false
	}
	if !isCloseEnough(p1.Z, p2.Z) {
		return false
	}
	return true
}

func TestTriangleHit0(t *testing.T) {
	someTriangle := Triangle{
		V0: vec3.Point{
			X: -1,
			Y: -1,
			Z: 0,
		},
		V1: vec3.Point{
			X: 1,
			Y: -1,
			Z: 0,
		},
		V2: vec3.Point{
			X: 0,
			Y: 1,
			Z: 0,
		},
		Mat: Lambertian{
			Albedo: vec3.Color{
				X: 0.5,
				Y: 0.0,
				Z: 0.0,
			},
		},
	}
	someTriangle.ComputeEdgesNormal()

	someRay := ray.Ray{
		Origin: vec3.Point{
			X: 0,
			Y: 0,
			Z: -1,
		},
		Direction: vec3.Point{
			X: 0,
			Y: 0,
			Z: 1,
		},
	}
	expectedHit := vec3.Point{
		X: 0,
		Y: 0,
		Z: 0,
	}
	expectedTime := 1.0
	rec := new(HitRecord)
	tmin := 0.001
	tmax := math.Inf(1)

	if actual := someTriangle.Hit(someRay, tmin, tmax, rec); actual {
		if !isCloseEnough(rec.T, expectedTime) {
			t.Errorf("Time when ray hits triangle is incorrect: expected=%f actual=%f", expectedTime, rec.T)
		}
		if !pointsEqual(rec.P, expectedHit) {
			t.Errorf("Ray hit triangle at incorrect point! expected=%#v, actual=%#v", expectedHit, rec.P)
		}
	} else {
		t.Errorf("Ray did not hit the triangle! This is wrong. Hit record: %#v", rec)
	}
}

func TestTriangleHit1(t *testing.T) {
	someTriangle := Triangle{
		V0: vec3.Point{
			X: -1,
			Y: -1,
			Z: 1,
		},
		V1: vec3.Point{
			X: 1,
			Y: -1,
			Z: 1,
		},
		V2: vec3.Point{
			X: 0,
			Y: 1,
			Z: 1,
		},
		Mat: Lambertian{
			Albedo: vec3.Color{
				X: 0.5,
				Y: 0.0,
				Z: 0.0,
			},
		},
	}
	someTriangle.ComputeEdgesNormal()

	someRay := ray.Ray{
		Origin: vec3.Point{
			X: 0,
			Y: 0,
			Z: 0,
		},
		Direction: vec3.Point{
			X: 0,
			Y: 0,
			Z: 1,
		},
	}
	expectedHit := vec3.Point{
		X: 0,
		Y: 0,
		Z: 1,
	}
	expectedTime := 1.0
	rec := new(HitRecord)
	tmin := 0.001
	tmax := math.Inf(1)

	if actual := someTriangle.Hit(someRay, tmin, tmax, rec); actual {
		if !isCloseEnough(rec.T, expectedTime) {
			t.Errorf("Time when ray hits triangle is incorrect: expected=%f actual=%f", expectedTime, rec.T)
		}
		if !pointsEqual(rec.P, expectedHit) {
			t.Errorf("Ray hit triangle at incorrect point! expected=%#v, actual=%#v", expectedHit, rec.P)
		}
	} else {
		t.Errorf("Ray did not hit the triangle! This is wrong. Hit record: %#v", rec)
	}
}

func TestTriangleHit2(t *testing.T) {
	someTriangle := Triangle{
		V0: vec3.Point{
			X: -1,
			Y: -1,
			Z: -1,
		},
		V1: vec3.Point{
			X: 1,
			Y: -1,
			Z: -1,
		},
		V2: vec3.Point{
			X: 0,
			Y: 1,
			Z: -1,
		},
		Mat: Lambertian{
			Albedo: vec3.Color{
				X: 0.5,
				Y: 0.0,
				Z: 0.0,
			},
		},
	}
	someTriangle.ComputeEdgesNormal()

	someRay := ray.Ray{
		Origin: vec3.Point{
			X: 0,
			Y: 0,
			Z: 0,
		},
		Direction: vec3.Point{
			X: 0,
			Y: 0,
			Z: -1,
		},
	}
	expectedHit := vec3.Point{
		X: 0,
		Y: 0,
		Z: -1,
	}
	expectedTime := 1.0
	rec := new(HitRecord)
	tmin := 0.001
	tmax := math.Inf(1)

	if actual := someTriangle.Hit(someRay, tmin, tmax, rec); actual {
		if !isCloseEnough(rec.T, expectedTime) {
			t.Errorf("Time when ray hits triangle is incorrect: expected=%f actual=%f", expectedTime, rec.T)
		}
		if !pointsEqual(rec.P, expectedHit) {
			t.Errorf("Ray hit triangle at incorrect point! expected=%#v, actual=%#v", expectedHit, rec.P)
		}
	} else {
		t.Errorf("Ray did not hit the triangle! This is wrong. Hit record: %#v", rec)
	}
}

func TestTriangleHit3(t *testing.T) {
	someTriangle := Triangle{
		V0: vec3.Point{
			X: -1,
			Y: -1,
			Z: -2,
		},
		V1: vec3.Point{
			X: 1,
			Y: -1,
			Z: -2,
		},
		V2: vec3.Point{
			X: 0,
			Y: 1,
			Z: -2,
		},
		Mat: Lambertian{
			Albedo: vec3.Color{
				X: 0.5,
				Y: 0.0,
				Z: 0.0,
			},
		},
	}
	someTriangle.ComputeEdgesNormal()

	someRay := ray.Ray{
		Origin: vec3.Point{
			X: 0,
			Y: 0,
			Z: 2,
		},
		Direction: vec3.Point{
			X: 0,
			Y: 0,
			Z: -1,
		},
	}
	expectedHit := vec3.Point{
		X: 0,
		Y: 0,
		Z: -2,
	}
	expectedTime := 4.0
	rec := new(HitRecord)
	tmin := 0.001
	tmax := math.Inf(1)

	if actual := someTriangle.Hit(someRay, tmin, tmax, rec); actual {
		if !isCloseEnough(rec.T, expectedTime) {
			t.Errorf("Time when ray hits triangle is incorrect: expected=%f actual=%f", expectedTime, rec.T)
		}
		if !pointsEqual(rec.P, expectedHit) {
			t.Errorf("Ray hit triangle at incorrect point! expected=%#v, actual=%#v", expectedHit, rec.P)
		}
		// t.Errorf("rec= %#v\n", rec)
	} else {
		t.Errorf("Ray did not hit the triangle! This is wrong. Hit record: %#v", rec)
	}
}
