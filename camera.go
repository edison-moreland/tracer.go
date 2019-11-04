package main

import "C"
import "math"

type Camera struct {
	Origin, lowerLeft, horizontal, vertical, u, v, w Vec3
	lensRadius                                       float64
}

type CameraOptions struct {
	LookFrom, LookAt, Up                 Vec3
	Fov, Aspect, Aperture, FocusDistance float64
}

func (c *Camera) Ray(s, t float64) Ray {
	rd := RandVec3InUnitDisk().Mul(c.lensRadius)
	offset := c.u.Mul(rd.X).Add(c.v.Mul(rd.Y))
	return Ray{
		Origin:    c.Origin.Add(offset),
		Direction: c.lowerLeft.Add(c.horizontal.Mul(s)).Add(c.vertical.Mul(t)).Sub(c.Origin).Sub(offset),
	}
}

func NewCamera(opts CameraOptions) Camera {
	theta := opts.Fov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := opts.Aspect * halfHeight

	w := opts.LookFrom.Sub(opts.LookAt).AsUnitVector()
	u := opts.Up.Cross(w).AsUnitVector()
	v := w.Cross(u)

	return Camera{
		Origin:     opts.LookFrom,
		lowerLeft:  opts.LookFrom.Sub(u.Mul(halfWidth * opts.FocusDistance)).Sub(v.Mul(halfHeight * opts.FocusDistance)).Sub(w.Mul(opts.FocusDistance)),
		horizontal: u.Mul(2 * halfWidth * opts.FocusDistance),
		vertical:   v.Mul(2 * halfHeight * opts.FocusDistance),
		lensRadius: opts.Aperture / 2,
		w:          w,
		u:          u,
		v:          v,
	}
}
