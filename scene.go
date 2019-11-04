package main

import "math/rand"

func RandomScene() HittableSlice {
	// World setup
	world := NewHittableSlice(
		// A place for spheres to sit
		&Sphere{
			Center:   Vec3{0.0, -1000, 0.0},
			Radius:   1000,
			Material: &Lambertian{Albedo: Vec3{0.5, 0.5, 0.5}},
		},

		// Show off materials
		&Sphere{
			Center:   Vec3{0.0, 1.0, 0.0},
			Radius:   1.0,
			Material: &Dielectric{RefractiveIndex: 1.5},
		},
		&Sphere{
			Center:   Vec3{-4.0, 1.0, 0.0},
			Radius:   1.0,
			Material: &Lambertian{Albedo: Vec3{0.4, 0.2, 0.1}},
		},
		&Sphere{
			Center:   Vec3{4.0, 1.0, 0.0},
			Radius:   1.0,
			Material: &Metal{Albedo: Vec3{0.7, 0.6, 0.5}, Diffusion: 0.0},
		},
	)

	// Make a LOT of spheres
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			// Choose random material
			materialChoice := rand.Float64()
			var material Material
			if materialChoice < 0.8 {
				material = &Lambertian{Albedo: Vec3{
					X: rand.Float64() * rand.Float64(),
					Y: rand.Float64() * rand.Float64(),
					Z: rand.Float64() * rand.Float64(),
				}}
			} else if materialChoice < 0.95 {
				material = &Metal{
					Albedo: Vec3{
						X: 0.5 * (1 + rand.Float64()),
						Y: 0.5 * (1 + rand.Float64()),
						Z: 0.5 * (1 + rand.Float64()),
					},
					Diffusion: 0.5 * rand.Float64(),
				}
			} else {
				material = &Dielectric{RefractiveIndex: 1.5}
			}

			world.AddHittable(&Sphere{
				Center:   Vec3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()},
				Radius:   0.2,
				Material: material,
			})
		}
	}

	return world
}
