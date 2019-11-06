package tracer

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/go-gl/mathgl/mgl64"

	"github.com/disintegration/imaging"
)

type ImageOptions struct {
	Width, Height int
	Path          string
}

type Image struct {
	ImageOptions
	*image.RGBA
}

func NewImage(options ImageOptions) Image {
	return Image{
		ImageOptions: options,
		RGBA:         image.NewRGBA(image.Rect(0, 0, options.Width, options.Height)),
	}
}

func RGBASetVec3(i *image.RGBA, vecColor mgl64.Vec3, x, y int) {
	// Map 0-1 to 0-MaxUint8 and XYZ to RGB
	max := float64(math.MaxUint8)
	i.Set(x, y, color.RGBA{
		R: uint8(max * vecColor[0]),
		G: uint8(max * vecColor[1]),
		B: uint8(max * vecColor[2]),
		A: 0xff,
	})
}

func (i *Image) Export() error {
	// Flip image, things are rendered upside down
	flipped := imaging.FlipV(i)

	file, err := os.Create(i.Path)
	if err != nil {
		return err
	}

	err = png.Encode(file, flipped)
	if err != nil {
		return err
	}

	return nil
}
