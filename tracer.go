package main

import (
	"fmt"
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
	nx := 180 * 2 // Image X
	ny := 135 * 2 // Image Y
	ns := 100     // Samples
	nb := 20      // Bounces
	out, err := NewPPM(uint(nx), uint(ny), "traced.ppm")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Camera setup
	lookFrom := Vec3{13, 2, 3}
	lookAt := Vec3{0.0, 0.0, 0.0}
	focusDistance := float64(10)
	options := CameraOptions{
		LookFrom:      lookFrom,
		LookAt:        lookAt,
		Up:            Vec3{0, 1, 0},
		Fov:           25,
		Aspect:        float64(nx) / float64(ny),
		Aperture:      0.05,
		FocusDistance: focusDistance,
	}
	camera := NewCamera(options)

	// World setup
	println("Generating scene...")
	world := RandomScene()

	// Render!
	println("Rendering...")
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
				fmt.Printf("Broke on pixel (%v, %v)", i, j)
				panic(err)
			}
		}
	}
}
