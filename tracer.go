package main

import (
	"image"
)

func main() {
	// Image setup
	img := NewImage(180, 135)

	// Camera setup
	options := RenderOptions{
		CameraOptions: CameraOptions{
			LookFrom:      Vec3{13, 2, 3},
			LookAt:        Vec3{0.0, 0.0, 0.0},
			Up:            Vec3{0, 1, 0},
			Fov:           25,
			Aspect:        float64(img.Rect.Max.X) / float64(img.Rect.Max.Y),
			Aperture:      0.05,
			FocusDistance: 10,
		},
		Samples: 10,
		Bounces: 20,
	}
	world := RandomWorld()
	scene := NewScene(options, &world)

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
