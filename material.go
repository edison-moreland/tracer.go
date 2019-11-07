package tracer

import (
	"math"
	"math/rand"

	"github.com/go-gl/mathgl/mgl64"
)

type Bounce struct {
	Attenuation mgl64.Vec3
	Scattered   Ray
}

type Material interface {
	Scatter(ray Ray, rec *HitRecord) *Bounce
}

type Lambertian struct {
	Albedo mgl64.Vec3
}

func (l Lambertian) Scatter(ray Ray, rec *HitRecord) *Bounce {
	target := rec.P.Add(rec.Normal).Add(RandVec3InUnitSphere())
	return &Bounce{
		Attenuation: l.Albedo,
		Scattered: Ray{
			Origin:    rec.P,
			Direction: target.Sub(rec.P),
		},
	}
}

type Metal struct {
	Albedo    mgl64.Vec3
	Diffusion float64 // 0.0 to 1.0
}

func (m Metal) Scatter(ray Ray, rec *HitRecord) *Bounce {
	reflected := Reflect(ray.Direction.Normalize(), rec.Normal)
	if reflected.Dot(rec.Normal) < 0 {
		return nil
	}

	// Add diffusion by randomizing direction
	direction := reflected.Add(RandVec3InUnitSphere().Mul(m.Diffusion))

	return &Bounce{
		Attenuation: m.Albedo,
		Scattered: Ray{
			Origin:    rec.P,
			Direction: direction,
		},
	}

}

func schlick(cosine, refractiveIndex float64) float64 {
	r0 := (1 - refractiveIndex) / (1 + refractiveIndex)
	r0 = r0 * r0
	r0 = r0 + (1-r0)*math.Pow(1-cosine, 5)
	return r0
}

type Dielectric struct {
	RefractiveIndex float64
}

func (d Dielectric) Scatter(ray Ray, rec *HitRecord) *Bounce {
	var outwardNormal mgl64.Vec3
	var niOverNt float64
	var cosine float64
	if ray.Direction.Dot(rec.Normal) > 0 {
		outwardNormal = Inverse(rec.Normal)
		niOverNt = d.RefractiveIndex
		cosine = d.RefractiveIndex * ray.Direction.Dot(rec.Normal) / ray.Direction.Len()
	} else {
		outwardNormal = rec.Normal
		niOverNt = 1.0 / d.RefractiveIndex
		cosine = -(ray.Direction.Dot(rec.Normal) / ray.Direction.Len())
	}

	direction := Reflect(ray.Direction, rec.Normal)
	if refracted, ok := Refract(ray.Direction, outwardNormal, niOverNt); ok {
		reflectProbability := schlick(cosine, d.RefractiveIndex)
		if rand.Float64() > reflectProbability {
			direction = refracted
		}
	}

	return &Bounce{
		Attenuation: mgl64.Vec3{1.0, 1.0, 1.0},
		Scattered:   Ray{rec.P, direction},
	}
}
