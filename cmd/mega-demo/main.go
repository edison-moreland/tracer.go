package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/edison-moreland/tracer.go"
	"github.com/edison-moreland/tracer.go/internal"
)

func main() {
	// Render Setup
	println("Generating scene...")

	scene := internal.GenDemoScene(1920, 1080, 200, 20, "traced.png")
	img := tracer.NewImage(scene.ImageOptions)

	fmt.Printf("Will render %vx%v image with %v samples per pixel\n", scene.Width, scene.Height, scene.Samples)

	// Start the render!
	n := runtime.NumCPU()
	fmt.Printf("Starting render on %v threads...\n", n)

	start := time.Now()
	scene.RenderParallel(img.RGBA, n)
	elapsed := time.Since(start)

	fmt.Printf("Rendered image in %v seconds\n", elapsed.Seconds())

	err := img.Export()
	if err != nil {
		println("Failed to export image")
		panic(err)
	}

	fmt.Printf("Exported image as %v", scene.Path)
}
