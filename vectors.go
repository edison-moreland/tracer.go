package tracer

import (
	"math"
	"math/rand"

	"github.com/go-gl/mathgl/mgl64"
)

func RandVec3InUnitSphere() (p mgl64.Vec3) {
	for {
		p = mgl64.Vec3{rand.Float64(), rand.Float64(), rand.Float64()}.Mul(2.0).Sub(mgl64.Vec3{1, 1, 1})
		if p.LenSqr() >= 1.0 {
			break
		}
	}
	return p
}

func RandVec3InUnitDisk() (p mgl64.Vec3) {
	for {
		p = mgl64.Vec3{rand.Float64(), rand.Float64(), 0}.Mul(2.0).Sub(mgl64.Vec3{1, 1, 0})
		if p.LenSqr() >= 1.0 {
			break
		}
	}
	return p
}

func Reflect(v mgl64.Vec3, n mgl64.Vec3) mgl64.Vec3 {
	return v.Sub(n.Mul(v.Dot(n) * 2))
}

func Refract(v mgl64.Vec3, n mgl64.Vec3, niOverNt float64) (refracted mgl64.Vec3, ok bool) {
	// Tries to refract across n, ok=false if cant
	uv := v.Normalize()
	dt := uv.Dot(n)
	discriminant := 1.0 - niOverNt*niOverNt*(1-dt*dt)
	if discriminant > 0 {
		refracted = uv.Sub(n.Mul(dt)).Mul(niOverNt).Sub(n.Mul(math.Sqrt(discriminant)))
		return refracted, true
	}
	return mgl64.Vec3{}, false
}

func Inverse(v mgl64.Vec3) mgl64.Vec3 {
	return v.Mul(-1)
}

func Div(v mgl64.Vec3, t float64) mgl64.Vec3 {
	return mgl64.Vec3{v[0] / t, v[1] / t, v[2] / t}
}

func MulByVec(v mgl64.Vec3, v2 mgl64.Vec3) mgl64.Vec3 {
	return mgl64.Vec3{v[0] * v2[0], v[1] * v2[1], v[2] * v2[2]}
}
