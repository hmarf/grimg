package rimg

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"

	"github.com/nfnt/resize"
)

type GifService struct {
	Img *gif.GIF
}

func keys(m map[color.Color]bool) color.Palette {
	var p color.Palette
	for k := range m {
		p = append(p, k)
	}
	return p
}

func (g *GifService) Resize(width uint, height uint) error {

	for i, frame := range g.Img.Image {
		rec := frame.Bounds()
		rImage := resize.Resize(width, height, frame.SubImage(rec), resize.Lanczos3)

		// The color used in the original image
		cUsedM := make(map[color.Color]bool)
		for x := 1; x <= rec.Dx(); x++ {
			for y := 1; y <= rec.Dy(); y++ {
				if _, ok := cUsedM[frame.At(x, y)]; !ok {
					cUsedM[frame.At(x, y)] = true
				}
			}
		}
		// scUsedP := keys(cUsedM)
		rrec := rImage.Bounds()
		if i > 0 {
			marginX := int(math.Floor(float64(rect.Min.X) * ratio))
			marginY := int(math.Floor(float64(rect.Min.Y) * ratio))
			resizedBounds = image.Rect(marginX, marginY, resizedBounds.Dx()+marginX,
				resizedBounds.Dy()+marginY)
		}
	}

	fmt.Println("ok")
	return nil
}
