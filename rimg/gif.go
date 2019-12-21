package rimg

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"os"

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

func (g *GifService) Resize(width uint, height uint, rate float64) error {

	// var out image.Image

	outGif := gif.GIF{}
	for i, frame := range g.Img.Image {
		rec := frame.Bounds()
		rImage := resize.Resize(width, height, frame.SubImage(rec), resize.Lanczos3)
		f, _ := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)
		defer f.Close()
		png.Encode(f, rImage)

		// The color used in the original image
		cUsedM := make(map[color.Color]bool)
		for x := 1; x <= rec.Dx(); x++ {
			for y := 1; y <= rec.Dy(); y++ {
				if _, ok := cUsedM[frame.At(x, y)]; !ok {
					cUsedM[frame.At(x, y)] = true
				}
			}
		}
		scUsedP := keys(cUsedM)
		rrec := rImage.Bounds()
		if i > 0 {
			x := int(width)
			y := int(height)
			rrec = image.Rect(0, 0, x, y)
		}
		rp := image.NewPaletted(rrec, scUsedP)
		draw.Draw(rp, rrec, rImage, image.ZP, draw.Src)

		outGif.Image = append(outGif.Image, rp)
		outGif.Delay = append(outGif.Delay, 100)
	}

	outGif.Config.Width = int(width)
	outGif.Config.Height = int(height)

	out, err := os.Create("resized.gif")
	if err != nil {
		return err
	}
	defer out.Close()

	err = gif.EncodeAll(out, &outGif)
	fmt.Println(err)
	fmt.Println("ok")
	return nil
}
