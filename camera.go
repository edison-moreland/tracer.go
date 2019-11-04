package main

import "C"
import "math"

type Camera struct {
	Origin, lowerLeft, horizontal, vertical, u, v, w Vec3
	lensRadius                                       float64
}

func (c *Camera) Ray(s, t float64) Ray {
	rd := RandVec3InUnitDisk().Mul(c.lensRadius)
	offset := c.u.Mul(rd.X).Add(c.v.Mul(rd.Y))
	return Ray{
		Origin:    c.Origin.Add(offset),
		Direction: c.lowerLeft.Add(c.horizontal.Mul(s)).Add(c.vertical.Mul(t)).Sub(c.Origin).Sub(offset),
	}
}

func NewCamera(lookFrom, lookAt, vup Vec3, fov, aspect, aperture, focusDistance float64) Camera {
	theta := fov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	w := lookFrom.Sub(lookAt).AsUnitVector()
	u := vup.Cross(w).AsUnitVector()
	v := w.Cross(u)

	return Camera{
		Origin:     lookFrom,
		lowerLeft:  lookFrom.Sub(u.Mul(halfWidth * focusDistance)).Sub(v.Mul(halfHeight * focusDistance)).Sub(w.Mul(focusDistance)),
		horizontal: u.Mul(2 * halfWidth * focusDistance),
		vertical:   v.Mul(2 * halfHeight * focusDistance),
		lensRadius: aperture / 2,
		w:          w, u: u, v: v,
	}
}
