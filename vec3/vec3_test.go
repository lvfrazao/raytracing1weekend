package vec3

import (
	"math"
	"testing"
)

func isCloseEnough(a, b float64) bool {
	tolerance := 0.00001
	return math.Abs(a-b) < tolerance
}

func TestIsCloseEnough1(t *testing.T) {
	a := 1.0001
	b := 1.0
	if isCloseEnough(a, b) {
		t.Errorf("%f and %f arent close enough, this should have failed", a, b)
	}
}

func TestIsCloseEnough2(t *testing.T) {
	a := 1.000001
	b := 1.0
	if !isCloseEnough(a, b) {
		t.Errorf("%f and %f are close enough, this should not have failed", a, b)
	}
}

func TestLenSqASM1(t *testing.T) {
	vector := Vec3{
		X: -10.0,
		Y: 102.5,
		Z: 11.0,
	}
	actualLen := vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z
	if vlen := lensq(vector.X, vector.Y, vector.Z); !isCloseEnough(vlen, actualLen) {
		t.Errorf("vector length incorrect, %f != %f", vlen, actualLen)
	}
}

func BenchmarkLenSqASM1(b *testing.B) {
	vector := Vec3{
		X: -10.0,
		Y: 102.5,
		Z: 11.0,
	}
	for n := 0; n < b.N; n++ {
		lensq(vector.X, vector.Y, vector.Z)
	}
}

func TestLenSqASM2(t *testing.T) {
	vector := Vec3{
		X: -10.0,
		Y: 102.5,
		Z: 11.0,
	}
	actualLen := vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z
	if vlen := lensq2(vector.X, vector.Y, vector.Z); !isCloseEnough(vlen, actualLen) {
		t.Errorf("vector length incorrect, %f != %f", vlen, actualLen)
	}
}

func BenchmarkLenSqASM2(b *testing.B) {
	vector := Vec3{
		X: -10.0,
		Y: 102.5,
		Z: 11.0,
	}
	for n := 0; n < b.N; n++ {
		lensq2(vector.X, vector.Y, vector.Z)
	}
}

func TestLenSqASM3(t *testing.T) {
	vector := Vec3{
		X: -10.0,
		Y: 102.5,
		Z: 11.0,
	}
	actualLen := vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z
	if vlen := lensq3(vector.X, vector.Y, vector.Z); !isCloseEnough(vlen, actualLen) {
		t.Errorf("vector length incorrect, %f != %f", vlen, actualLen)
	}
}

func BenchmarkLenSqASM3(b *testing.B) {
	vector := Vec3{
		X: -10.0,
		Y: 102.5,
		Z: 11.0,
	}
	for n := 0; n < b.N; n++ {
		lensq3(vector.X, vector.Y, vector.Z)
	}
}

func TestLenSqASM4(t *testing.T) {
	vector := Vec3{
		X: -10.0,
		Y: 102.5,
		Z: 11.0,
	}
	actualLen := vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z
	if vlen := lensq4(vector.X, vector.Y, vector.Z); !isCloseEnough(vlen, actualLen) {
		t.Errorf("vector length incorrect, %f != %f", vlen, actualLen)
	}
}

func BenchmarkLenSqASM4(b *testing.B) {
	vector := Vec3{
		X: -10.0,
		Y: 102.5,
		Z: 11.0,
	}
	for n := 0; n < b.N; n++ {
		lensq4(vector.X, vector.Y, vector.Z)
	}
}

func TestLenSqASM5(t *testing.T) {
	vector := Vec3{
		X: -10.0,
		Y: 102.5,
		Z: 11.0,
	}
	actualLen := vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z
	if vlen := lensq5(vector.X, vector.Y, vector.Z); !isCloseEnough(vlen, actualLen) {
		t.Errorf("vector length incorrect, %f != %f", vlen, actualLen)
	}
}

func BenchmarkLenSqASM5(b *testing.B) {
	vector := Vec3{
		X: -10.0,
		Y: 102.5,
		Z: 11.0,
	}
	for n := 0; n < b.N; n++ {
		lensq5(vector.X, vector.Y, vector.Z)
	}
}

func TestLenSqGo1(t *testing.T) {
	vector := Vec3{
		X: -10.0,
		Y: 102.5,
		Z: 11.0,
	}
	actualLen := vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z
	if vlen := vector.Dot(vector); !isCloseEnough(vlen, actualLen) {
		t.Errorf("vector length incorrect, %f != %f", vlen, actualLen)
	}
}

func BenchmarkLenSqGo1(b *testing.B) {
	vector := Vec3{
		X: -10.0,
		Y: 102.5,
		Z: 11.0,
	}
	for n := 0; n < b.N; n++ {
		vector.Dot(vector)
	}
}

func TestLengthSquared(t *testing.T) {
	vector := Vec3{
		X: -10.0,
		Y: 102.5,
		Z: 11.0,
	}
	actualLen := vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z
	if vlen := vector.LengthSquared(); !isCloseEnough(vlen, actualLen) {
		t.Errorf("vector length incorrect, %f != %f", vlen, actualLen)
	}
}

func BenchmarkLenghtSquared(b *testing.B) {
	vector := Vec3{
		X: -10.0,
		Y: 102.5,
		Z: 11.0,
	}
	for n := 0; n < b.N; n++ {
		vector.LengthSquared()
	}
}
