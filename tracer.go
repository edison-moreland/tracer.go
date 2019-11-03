package main

import (
	"math"
	"math/rand"
)

func LinearInterpolation(t float64, color1, color2 Vec3) Vec3 {
	return color1.Mul(1.0 - t).Add(color2.Mul(t))
}

func BackgroundColor(ray Ray) Vec3 {
	unitDirection := ray.Direction.AsUnitVector()
	t := 0.5 * (unitDirection.Y + 1.0)
	return LinearInterpolation(t, Vec3{1.0, 1.0, 1.0}, Vec3{0.5, 0.7, 1.0})
}

func Trace(ray Ray, world Hittable, bounces int) (color Vec3) {
	if bounces <= 0 {
		return Vec3{}
	}

	if rec := world.Hit(ray, 0.001, math.MaxFloat64); rec != nil {
		if bounce := rec.Material.Scatter(ray, rec); bounce != nil {
			return Trace(bounce.Scattered, world, bounces-1).MulByVec(bounce.Attenuation)
		}
		return Vec3{}
	}

	return BackgroundColor(ray)
}

func main() {
	// Image setup
	nx := 200 // Image X
	ny := 100 // Image Y
	ns := 100 // Samples
	nb := 20  // Bounces
	out, err := NewPPM(uint(nx), uint(ny), "traced.ppm")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Camera setup
	camera := DefaultCamera()

	// World setup
	world := NewHittableSlice(
		&Sphere{
			Center:   Vec3{0.0, 0.0, -1.0},
			Radius:   0.5,
			Material: &Lambertian{Albedo: Vec3{0.1, 0.2, 0.5}},
		},
		&Sphere{
			Center:   Vec3{0.0, -100.5, -1.0},
			Radius:   100,
			Material: &Lambertian{Albedo: Vec3{0.8, 0.8, 0.0}},
		},
		&Sphere{
			Center:   Vec3{1.0, 0.0, -1.0},
			Radius:   0.5,
			Material: &Metal{Albedo: Vec3{0.8, 0.6, 0.2}, Diffusion: 0.0},
		},
		&Sphere{
			Center:   Vec3{-1.0, 0.0, -1.0},
			Radius:   0.5,
			Material: &Dielectric{RefractiveIndex: 1},
		},
	)

	// Render!
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			// Sample pixel
			var color Vec3
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				camRay := camera.Ray(u, v)
				color = color.Add(Trace(camRay, &world, nb))
			}
			color = color.Div(float64(ns))

			// Gamma magic (Brightens image)
			color = Vec3{
				X: math.Sqrt(color.X),
				Y: math.Sqrt(color.Y),
				Z: math.Sqrt(color.Z),
			}

			// Write pixel
			err = out.WriteVec3(color)
			if err != nil {
				panic(err)
			}
		}
	}
}
