package image

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

type PngService struct {
	img *image.Image
}

func (p *PngService) Resize(width uint, height uint) error {
	m := resize.Resize(width, height, *p.img, resize.Lanczos3)

	outfile, err := os.Create("resized.jpg")
	if err != nil {
		return err
	}
	defer outfile.Close()

	return jpeg.Encode(outfile, m, &jpeg.Options{Quality: 100})
}
