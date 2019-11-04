package main

import (
	"math"
)

func main() {
	// Image setup
	nx := 180 // Image X
	ny := 135 // Image Y
	out, err := NewImage(nx, ny)
	if err != nil {
		panic(err)
	}

	// Camera setup
	options := RenderOptions{
		CameraOptions: CameraOptions{
			LookFrom:      Vec3{13, 2, 3},
			LookAt:        Vec3{0.0, 0.0, 0.0},
			Up:            Vec3{0, 1, 0},
			Fov:           25,
			Aspect:        float64(nx) / float64(ny),
			Aperture:      0.05,
			FocusDistance: 10,
		},
		Samples: 10,
		Bounces: 20,
	}
	world := RandomWorld()
	scene := NewScene(options, &world)

	// Render!
	for j := 0; j < ny; j++ {
		for i := 0; i < nx; i++ {
			// Sample pixel
			color := scene.SamplePixel(float64(i), float64(j), float64(nx), float64(ny))

			// Gamma magic (Brightens image)
			color = Vec3{
				X: math.Sqrt(color.X),
				Y: math.Sqrt(color.Y),
				Z: math.Sqrt(color.Z),
			}

			// Write pixel
			out.SetVec3(color, i, j)
		}
	}

	err = out.ExportPNG("traced.png")
	if err != nil {
		println("Failed to export image")
		panic(err)
	}
}
