package rimg

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

type jpegService struct {
	Img *image.Image
}

func (j *jpegService) resize(width uint, height uint) error {
	m := resize.Resize(width, height, *j.Img, resize.Lanczos3)

	outfile, err := os.Create("resize.jpg")
	if err != nil {
		return err
	}
	defer outfile.Close()

	return jpeg.Encode(outfile, m, &jpeg.Options{Quality: 100})
}
