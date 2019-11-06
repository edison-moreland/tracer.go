package tracer

import (
	"image"
	"math"
	"math/rand"

	"github.com/go-gl/mathgl/mgl64"
)

func LinearInterpolation(t float64, color1, color2 mgl64.Vec3) mgl64.Vec3 {
	return color1.Mul(1.0 - t).Add(color2.Mul(t))
}

func BackgroundColor(ray Ray) mgl64.Vec3 {
	unitDirection := ray.Direction.Normalize()
	t := 0.5 * (unitDirection.Y() + 1.0)
	return LinearInterpolation(t, mgl64.Vec3{1.0, 1.0, 1.0}, mgl64.Vec3{0.5, 0.7, 1.0})
}

func Trace(ray Ray, world Primitive, bounces int) (color mgl64.Vec3) {
	if bounces <= 0 {
		return mgl64.Vec3{}
	}

	if rec := world.Hit(ray, mgl64.Epsilon, math.MaxFloat64); rec != nil {
		if bounce := rec.Material.Scatter(ray, rec); bounce != nil {
			return MulByVec(Trace(bounce.Scattered, world, bounces-1), bounce.Attenuation)
		}
		return mgl64.Vec3{}
	}

	return BackgroundColor(ray)
}

type RenderOptions struct {
	CameraOptions
	ImageOptions
	Samples, Bounces int
}

// Scene describes everything needed to render and image
type Scene struct {
	RenderOptions
	camera Camera
	world  Primitive
}

func NewScene(options RenderOptions, world Primitive) Scene {
	return Scene{
		RenderOptions: options,
		camera:        NewCamera(options.CameraOptions),
		world:         world,
	}
}

func (s *Scene) SamplePixel(x, y float64) mgl64.Vec3 {
	// Average multiple samples with random offsets
	var color mgl64.Vec3
	for i := 0; i < s.Samples; i++ {
		// Translate pixel to camera plane
		u := (x + rand.Float64()) / float64(s.Width)
		v := (y + rand.Float64()) / float64(s.Height)

		// New ray from camera origin to point on camera plane
		camRay := s.camera.Ray(u, v)

		// Bounce around
		color = color.Add(Trace(camRay, s.world, s.Bounces))
	}
	return Div(color, float64(s.Samples))
}

func (s *Scene) RenderToRGBA(img *image.RGBA) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Sample pixel
			color := s.SamplePixel(float64(x), float64(y))

			// Gamma magic (Brightens image)
			color = mgl64.Vec3{
				math.Sqrt(color[0]),
				math.Sqrt(color[1]),
				math.Sqrt(color[2]),
			}

			// Write pixel
			RGBASetVec3(img, color, x, y)
		}
	}
}
