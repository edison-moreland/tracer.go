package main

import (
	"fmt"
	"math"
	"os"
)

func LinearInterpolation(t float64, color1, color2 Vec3) Vec3 {
	return color1.Mul(1.0 - t).Add(color2.Mul(t))
}

func BackgroundColor(ray Ray) Vec3 {
	unitDirection := ray.Direction.AsUnitVector()
	t := 0.5 * (unitDirection.Y + 1.0)
	return LinearInterpolation(t, NewVec3(1.0, 1.0, 1.0), NewVec3(0.5, 0.7, 1.0))
}

func main() {
	out, err := os.Create("traced.ppm")
	if err != nil {
		panic(err.Error())
	}
	defer out.Close()

	// Image setup
	nx := 200
	ny := 100
	_, _ = fmt.Fprintf(out, "P3\n%d %d\n255\n", nx, ny)

	// Camera setup
	lowerLeftCorner := NewVec3(-2.0, -1.0, -1.0)
	horizontal := NewVec3(4.0, 0.0, 0.0)
	vertical := NewVec3(0.0, 2.0, 0.0)
	origin := NewVec3(0.0, 0.0, 0.0)

	// Scene setup
	scene := NewHittableSlice(
		&Sphere{
			Center: NewVec3(0.0, 0.0, -1.0),
			Radius: 0.5,
		},
		&Sphere{
			Center: NewVec3(0.0, -100.5, -1.0),
			Radius: 100,
		},
	)

	// Render!
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			// Construct camera
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			camRay := Ray{
				Origin:    origin,
				Direction: lowerLeftCorner.Add(horizontal.Mul(u)).Add(vertical.Mul(v)),
			}

			color := BackgroundColor(camRay)
			if rec := scene.Hit(camRay, 0.0, math.MaxFloat64); rec != nil {
				color = NewVec3(rec.normal.X+1, rec.normal.Y+1, rec.normal.Z+1).Mul(0.5)
			}

			ir := int(255.99 * color.X)
			ig := int(255.99 * color.Y)
			ib := int(255.99 * color.Z)
			_, _ = fmt.Fprintf(out, "%d %d %d\n", ir, ig, ib)
		}
	}

}
