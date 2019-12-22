package rimg

import (
	"image"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

type pngService struct {
	Img *image.Image
}

func (p *pngService) resize(width uint, height uint, o Option) error {
	m := resize.Resize(width, height, *p.Img, resize.Lanczos3)

	outfile, err := os.Create(o.OutputFile)
	if err != nil {
		return err
	}
	defer outfile.Close()

	return png.Encode(outfile, m)
}
