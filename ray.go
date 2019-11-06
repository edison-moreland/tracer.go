package tracer

type Ray struct {
	Origin, Direction Vec3
}

func (r Ray) PointOnRay(t float64) Vec3 {
	return r.Origin.Add(r.Direction.Mul(t))
}
