package tracer

import "github.com/go-gl/mathgl/mgl64"

type Ray struct {
	Origin, Direction mgl64.Vec3
}

func (r Ray) PointOnRay(t float64) mgl64.Vec3 {
	return r.Origin.Add(r.Direction.Mul(t))
}
