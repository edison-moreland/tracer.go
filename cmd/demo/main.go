package main

import (
	"fmt"
	"image"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/go-gl/mathgl/mgl64"

	"github.com/edison-moreland/tracer.go"
)

func RandomWorld() tracer.Primitives {
	// World setup
	world := tracer.NewPrimitiveSlice(
		// A place for spheres to sit
		&tracer.Sphere{
			Center:   mgl64.Vec3{0.0, -1000, 0.0},
			Radius:   1000,
			Material: &tracer.Lambertian{Albedo: mgl64.Vec3{0.5, 0.5, 0.5}},
		},

		// Show off materials
		&tracer.Sphere{
			Center:   mgl64.Vec3{0.0, 1.0, 0.0},
			Radius:   1.0,
			Material: &tracer.Dielectric{RefractiveIndex: 1.5},
		},
		&tracer.Sphere{
			Center:   mgl64.Vec3{-4.0, 1.0, 0.0},
			Radius:   1.0,
			Material: &tracer.Lambertian{Albedo: mgl64.Vec3{0.4, 0.2, 0.1}},
		},
		&tracer.Sphere{
			Center:   mgl64.Vec3{4.0, 1.0, 0.0},
			Radius:   1.0,
			Material: &tracer.Metal{Albedo: mgl64.Vec3{0.7, 0.6, 0.5}, Diffusion: 0.0},
		},
	)

	// Make a LOT of spheres
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			// Choose random material
			materialChoice := rand.Float64()
			var material tracer.Material
			if materialChoice < 0.8 {
				material = &tracer.Lambertian{Albedo: mgl64.Vec3{
					rand.Float64() * rand.Float64(),
					rand.Float64() * rand.Float64(),
					rand.Float64() * rand.Float64(),
				}}
			} else if materialChoice < 0.95 {
				material = &tracer.Metal{
					Albedo: mgl64.Vec3{
						0.5 * (1 + rand.Float64()),
						0.5 * (1 + rand.Float64()),
						0.5 * (1 + rand.Float64()),
					},
					Diffusion: 0.5 * rand.Float64(),
				}
			} else {
				material = &tracer.Dielectric{RefractiveIndex: 1.5}
			}

			world.AddPrimitive(&tracer.Sphere{
				Center:   mgl64.Vec3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()},
				Radius:   0.2,
				Material: material,
			})
		}
	}

	return world
}

func ToColumns(rect image.Rectangle, n int) []image.Rectangle {
	x0, y0 := rect.Min.X, rect.Min.Y
	x1, y1 := rect.Max.X, rect.Max.Y

	rects := make([]image.Rectangle, n)
	for i := 0; i < n; i++ {
		rects[i] = image.Rect(x0+(x1/n)*i, y0, x0+(x1/n)*(i+1), y1)
	}

	return rects
}

func main() {
	// Image setup
	img := tracer.NewImage(180, 135)

	// Camera setup
	options := tracer.RenderOptions{
		CameraOptions: tracer.CameraOptions{
			LookFrom:      mgl64.Vec3{13, 2, 3},
			LookAt:        mgl64.Vec3{0, 0, 0},
			Up:            mgl64.Vec3{0, 1, 0},
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
	start := time.Now()
	n := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(n)

	for _, col := range ToColumns(img.Bounds(), n) {
		sub := img.SubImage(col).(*image.RGBA)

		// Render!
		go func() {
			scene.RenderToRGBA(sub, img.Width, img.Height)
			wg.Done()
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Print(elapsed.Seconds())

	err := img.ExportPNG("traced.png")
	if err != nil {
		println("Failed to export image")
		panic(err)
	}
}
