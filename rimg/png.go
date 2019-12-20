package rimg

import (
	"image"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

type PngService struct {
	Img *image.Image
}

func (p *PngService) Resize(width uint, height uint) error {
	m := resize.Resize(width, height, *p.Img, resize.Lanczos3)

	outfile, err := os.Create("resize.png")
	if err != nil {
		return err
	}
	defer outfile.Close()

	return png.Encode(outfile, m)
}
