package main

type Bounce struct {
	Attenuation Vec3
	Scattered   Ray
}

type Material interface {
	Scatter(ray Ray, rec *HitRecord) *Bounce
}

type Lambertian struct {
	Albedo Vec3
}

func (l *Lambertian) Scatter(ray Ray, rec *HitRecord) *Bounce {
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
	Albedo    Vec3
	Diffusion float64 // 0.0 to 1.0
}

func (m *Metal) Scatter(ray Ray, rec *HitRecord) *Bounce {
	reflected := ray.Direction.AsUnitVector().ReflectAcross(rec.Normal)
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
