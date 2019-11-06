package tracer

import (
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

type HitRecord struct {
	T         float64
	P, Normal mgl64.Vec3
	Material  Material
}

// Future improvement: rename to "object" or "geometry" or something else, hittable seems so awkward
type Hittable interface {
	// Future improvement: return *HitRecord instead of bool, nil return mean no hit found
	// Disadvantage: a new HitRecord is allocated every time a ray hits an object, instead of just once
	Hit(r Ray, tMin, tMax float64) *HitRecord
}

type Sphere struct {
	Center   mgl64.Vec3
	Radius   float64
	Material Material
}

func (s *Sphere) Hit(r Ray, tMin, tMax float64) *HitRecord {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.LenSqr()
	b := oc.Dot(r.Direction)
	c := oc.LenSqr() - s.Radius*s.Radius
	discriminant := b*b - a*c

	if discriminant < 0 {
		return nil
	}

	if t := (-b - math.Sqrt(discriminant)) / a; t < tMax && t > tMin {
		rec := &HitRecord{}
		rec.T = t
		rec.P = r.PointOnRay(t)
		rec.Normal = Div(rec.P.Sub(s.Center), s.Radius)
		rec.Material = s.Material
		return rec
	}

	if t := (-b + math.Sqrt(discriminant)) / a; t < tMax && t > tMin {
		rec := &HitRecord{}
		rec.T = t
		rec.P = r.PointOnRay(t)
		rec.Normal = Div(rec.P.Sub(s.Center), s.Radius)
		rec.Material = s.Material
		return rec
	}

	return nil
}

type HittableSlice struct {
	hittables []Hittable
}

func NewHittableSlice(hittables ...Hittable) HittableSlice {
	return HittableSlice{hittables: hittables}
}

func (h *HittableSlice) AddHittable(hittable Hittable) {
	h.hittables = append(h.hittables, hittable)
}

func (h *HittableSlice) AddHittables(hittables ...Hittable) {
	h.hittables = append(h.hittables, hittables...)
}

func (h *HittableSlice) Hit(r Ray, tMin, tMax float64) *HitRecord {
	var rec *HitRecord = nil
	closest := tMax
	for _, hittable := range h.hittables {
		if newRec := hittable.Hit(r, tMin, closest); newRec != nil {
			closest = newRec.T
			rec = newRec
		}
	}
	return rec
}
