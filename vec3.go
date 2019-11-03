package main

import (
	"math"
	"math/rand"
)

type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) Inverse() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

func RandVec3InUnitSphere() (p Vec3) {
	for {
		p = Vec3{rand.Float64(), rand.Float64(), rand.Float64()}.Mul(2.0).Sub(Vec3{1, 1, 1})
		if p.SquaredLength() >= 1.0 {
			break
		}
	}
	return p
}

func (v Vec3) Reflect(n Vec3) Vec3 {
	return v.Sub(n.Mul(v.Dot(n) * 2))
}

func (v Vec3) Refract(n Vec3, niOverNt float64) (refracted Vec3, ok bool) {
	// Tries to refract across n, ok=false if cant
	uv := v.AsUnitVector()
	dt := uv.Dot(n)
	discriminant := 1.0 - niOverNt*niOverNt*(1-dt*dt)
	if discriminant > 0 {
		refracted = uv.Sub(n.Mul(dt)).Mul(niOverNt).Sub(n.Mul(math.Sqrt(discriminant)))
		return refracted, true
	}
	return Vec3{}, false
}

func (v Vec3) Dot(v2 Vec3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{v.Y*v2.Z - v.Z*v2.Y, -(v.X*v2.Z - v.Z*v2.X), v.X*v2.Y - v.Y*v2.X}
}

func (v Vec3) SquaredLength() float64 {
	return v.Dot(v)
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.SquaredLength())
}

func (v Vec3) AsUnitVector() Vec3 {
	return v.Div(v.Length())
}

func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

func (v Vec3) Mul(t float64) Vec3 {
	return Vec3{v.X * t, v.Y * t, v.Z * t}
}

func (v Vec3) Div(t float64) Vec3 {
	return Vec3{v.X / t, v.Y / t, v.Z / t}
}

func (v Vec3) MulByVec(v2 Vec3) Vec3 {
	return Vec3{v.X * v2.X, v.Y * v2.Y, v.Z * v2.Z}
}

func (v Vec3) DivByVec(v2 Vec3) Vec3 {
	return Vec3{v.X / v2.X, v.Y / v2.Y, v.Z / v2.Z}
}
