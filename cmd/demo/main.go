package main

import (
	"image"
	"math/rand"

	"github.com/edison-moreland/tracer.go"
)

func RandomWorld() tracer.HittableSlice {
	// World setup
	world := tracer.NewHittableSlice(
		// A place for spheres to sit
		&tracer.Sphere{
			Center:   tracer.Vec3{0.0, -1000, 0.0},
			Radius:   1000,
			Material: &tracer.Lambertian{Albedo: tracer.Vec3{0.5, 0.5, 0.5}},
		},

		// Show off materials
		&tracer.Sphere{
			Center:   tracer.Vec3{0.0, 1.0, 0.0},
			Radius:   1.0,
			Material: &tracer.Dielectric{RefractiveIndex: 1.5},
		},
		&tracer.Sphere{
			Center:   tracer.Vec3{-4.0, 1.0, 0.0},
			Radius:   1.0,
			Material: &tracer.Lambertian{Albedo: tracer.Vec3{0.4, 0.2, 0.1}},
		},
		&tracer.Sphere{
			Center:   tracer.Vec3{4.0, 1.0, 0.0},
			Radius:   1.0,
			Material: &tracer.Metal{Albedo: tracer.Vec3{0.7, 0.6, 0.5}, Diffusion: 0.0},
		},
	)

	// Make a LOT of spheres
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			// Choose random material
			materialChoice := rand.Float64()
			var material tracer.Material
			if materialChoice < 0.8 {
				material = &tracer.Lambertian{Albedo: tracer.Vec3{
					X: rand.Float64() * rand.Float64(),
					Y: rand.Float64() * rand.Float64(),
					Z: rand.Float64() * rand.Float64(),
				}}
			} else if materialChoice < 0.95 {
				material = &tracer.Metal{
					Albedo: tracer.Vec3{
						X: 0.5 * (1 + rand.Float64()),
						Y: 0.5 * (1 + rand.Float64()),
						Z: 0.5 * (1 + rand.Float64()),
					},
					Diffusion: 0.5 * rand.Float64(),
				}
			} else {
				material = &tracer.Dielectric{RefractiveIndex: 1.5}
			}

			world.AddHittable(&tracer.Sphere{
				Center:   tracer.Vec3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()},
				Radius:   0.2,
				Material: material,
			})
		}
	}

	return world
}

func main() {
	// Image setup
	img := tracer.NewImage(180, 135)

	// Camera setup
	options := tracer.RenderOptions{
		CameraOptions: tracer.CameraOptions{
			LookFrom:      tracer.Vec3{X: 13, Y: 2, Z: 3},
			LookAt:        tracer.Vec3{X: 0, Y: 0, Z: 0},
			Up:            tracer.Vec3{X: 0, Y: 1, Z: 0},
			Fov:           25,
			Aspect:        float64(img.Rect.Max.X) / float64(img.Rect.Max.Y),
			Aperture:      0.05,
			FocusDistance: 10,
		},
		Samples: 10,
		Bounces: 20,
	}
	world := RandomWorld()
	scene := tracer.NewScene(options, &world)

	// Little test
	sub := img.SubImage(img.Bounds()).(*image.RGBA)

	// Render!
	scene.RenderToRGBA(sub)

	err := img.ExportPNG("traced.png")
	if err != nil {
		println("Failed to export image")
		panic(err)
	}
}
