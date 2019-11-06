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

type Image struct {
	*image.RGBA
	Width, Height int
}

func NewImage(width, height int) Image {
	return Image{
		RGBA:   image.NewRGBA(image.Rect(0, 0, width, height)),
		Width:  width,
		Height: height,
	}
}

func RGBASetVec3(i *image.RGBA, vecColor mgl64.Vec3, x, y int) {
	// Map 0-1 to 0-MaxUint8 and XYZ to RGB
	max := float64(math.MaxUint8)
	i.Set(x, y, color.RGBA{
		R: uint8(max * vecColor.X()),
		G: uint8(max * vecColor.Y()),
		B: uint8(max * vecColor.Z()),
		A: 0xff,
	})
}

func (i *Image) ExportPNG(path string) error {
	// Flip image, things are rendered upside down
	flipped := imaging.FlipV(i)

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	err = png.Encode(file, flipped)
	if err != nil {
		return err
	}

	return nil
}
