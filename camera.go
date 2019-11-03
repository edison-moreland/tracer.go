package main

type Camera struct {
	Origin, LowerLeft, Horizontal, Vertical Vec3
}

func DefaultCamera() Camera {
	return Camera{
		Origin:     Vec3{0.0, 0.0, 0.0},
		LowerLeft:  Vec3{-2.0, -1.0, -1.0},
		Horizontal: Vec3{4.0, 0.0, 0.0},
		Vertical:   Vec3{0.0, 2.0, 0.0},
	}
}

func (c *Camera) Ray(u, v float64) Ray {
	return Ray{
		Origin:    c.Origin,
		Direction: c.LowerLeft.Add(c.Horizontal.Mul(u)).Add(c.Vertical.Mul(v)),
	}
}
