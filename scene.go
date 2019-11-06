package tracer

import (
	"image"
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

type RenderOptions struct {
	CameraOptions
	Samples, Bounces int
}

// Scene describes everything needed to render and image
type Scene struct {
	RenderOptions
	camera Camera
	world  Hittable
}

func NewScene(options RenderOptions, world Hittable) Scene {
	return Scene{
		RenderOptions: options,
		camera:        NewCamera(options.CameraOptions),
		world:         world,
	}
}

func (s *Scene) SamplePixel(x, y, xMax, yMax float64) Vec3 {
	// Average multiple samples with random offsets
	var color Vec3
	for i := 0; i < s.Samples; i++ {
		// Translate pixel to camera plane
		u := (x + rand.Float64()) / xMax
		v := (y + rand.Float64()) / yMax

		// New ray from camera origin to point on camera plane
		camRay := s.camera.Ray(u, v)

		// Bounce around
		color = color.Add(Trace(camRay, s.world, s.Bounces))
	}
	return color.Div(float64(s.Samples))
}

func (s *Scene) RenderToRGBA(img *image.RGBA) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Sample pixel
			color := s.SamplePixel(float64(x), float64(y), float64(bounds.Max.X), float64(bounds.Max.Y))

			// Gamma magic (Brightens image)
			color = Vec3{
				X: math.Sqrt(color.X),
				Y: math.Sqrt(color.Y),
				Z: math.Sqrt(color.Z),
			}

			// Write pixel
			RGBASetVec3(img, color, x, y)
		}
	}
}
