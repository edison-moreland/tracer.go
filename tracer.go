package main

import (
	"fmt"
	"math"
	"os"
)

type Vec3 struct{
	X, Y, Z float64
}

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}

func (v Vec3) Inverse() Vec3 {
	return NewVec3(-v.X, -v.Y, -v.Z)
}

func (v Vec3) Dot(v2 Vec3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v Vec3) Cross(v2 Vec3) Vec3 {
	return NewVec3(v.Y*v2.Z - v.Z*v2.Y, -(v.X*v2.Z - v.Z*v2.X), v.X*v2.Y - v.Y*v2.X)
}

func (v Vec3) SquaredLength() float64 {
	return v.Dot(v)
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.SquaredLength())
}

func (v Vec3) AsUnitVector() Vec3 {
	k := 1.0 / v.Length()
	return NewVec3(v.X * k, v.Y * k, v.Z * k)
}


func (v Vec3) Add(v2 Vec3) Vec3 {
	return NewVec3(v.X+v2.X, v.Y+v2.Y, v.Z+v2.Z)
}

func (v Vec3) Sub(v2 Vec3) Vec3 {
	return NewVec3(v.X-v2.X, v.Y-v2.Y, v.Z-v2.Z)
}

func (v Vec3) Mul(t float64) Vec3 {
	return NewVec3(v.X*t, v.Y*t, v.Z*t)
}

func (v Vec3) Div(t float64) Vec3 {
	return NewVec3(v.X/t, v.Y/t, v.Z/t)
}

func (v Vec3) MulByVec(v2 Vec3) Vec3 {
	return NewVec3(v.X*v2.X, v.Y*v2.Y, v.Z*v2.Z)
}

func (v Vec3) DivByVec(v2 Vec3) Vec3 {
	return NewVec3(v.X/v2.X, v.Y/v2.Y, v.Z/v2.Z)
}


func main() {
	out, err := os.Create("traced.ppm")
	if err != nil { panic(err.Error()) }
	defer out.Close()

	nx := 200
	ny := 100
	_, _ = fmt.Fprintf(out, "P3\n%d %d\n255\n", nx, ny)
	for j := ny-1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			color := NewVec3(float64(i) / float64(nx), float64(j) / float64(ny), 0.2)
			ir := int(255.99*color.X)
			ig := int(255.99*color.Y)
			ib := int(255.99*color.Z)
			_, _ = fmt.Fprintf(out, "%d %d %d\n", ir, ig, ib)
		}
	}
	
}