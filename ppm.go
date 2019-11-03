package main

import (
	"fmt"
	"os"
)

type PPM struct {
	*os.File
	Width, Height uint
}

func NewPPM(width, height uint, path string) (PPM, error) {
	file, err := os.Create(path)
	if err != nil {
		return PPM{}, err
	}

	// Write PPM header
	_, err = fmt.Fprintf(file, "P3\n%d %d\n255\n", width, height)
	if err != nil {
		file.Close()
		return PPM{}, err
	}

	return PPM{
		File:   file,
		Width:  width,
		Height: height,
	}, nil
}

func (p *PPM) WriteVec3(color Vec3) error {
	return p.WriteColor(
		int(255.99*color.X),
		int(255.99*color.Y),
		int(255.99*color.Z),
	)
}

func (p *PPM) WriteColor(r, g, b int) error {
	if (r < 0 || g < 0 || b < 0) || (r > 255 || g > 255 || b > 255) {
		return fmt.Errorf("color out of range: (%v, %v, %v)", r, g, b)
	}

	_, err := fmt.Fprintf(p, "%d %d %d\n", r, g, b)
	if err != nil {
		return err
	}

	return nil
}
