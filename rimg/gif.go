package rimg

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"

	"github.com/nfnt/resize"
)

type gifService struct {
	Img *gif.GIF
}

func keys(m map[color.Color]bool) color.Palette {
	var p color.Palette
	for k := range m {
		p = append(p, k)
	}
	return p
}

func (g *gifService) resize(width uint, height uint, o Option) error {

	outGif := gif.GIF{}
	cUsedM := make(map[color.Color]bool)
	fmt.Println(width, height)
	var u int
	var e int
	for _, frame := range g.Img.Image {

		rec := frame.Bounds()

		rImage := resize.Resize(width, height, frame.SubImage(rec), resize.Lanczos3)
		rrec := rImage.Bounds()

		if rrec.Dx() > int(width) || rrec.Dy() > int(height) {
			e++
			continue
		}

		// The color used in the original image
		for x := 1; x <= rec.Dx(); x++ {
			for y := 1; y <= rec.Dy(); y++ {
				if _, ok := cUsedM[frame.At(x, y)]; !ok {
					cUsedM[frame.At(x, y)] = true
				}
			}
		}
		scUsedP := keys(cUsedM)
		if u > 0 {
			rrec = image.Rect(
				int(float64(rec.Min.X)*o.Compression),
				int(float64(rec.Min.Y)*o.Compression),
				rrec.Dx()+int(float64(rec.Min.X)*o.Compression),
				rrec.Dy()+int(float64(rec.Min.Y)*o.Compression))
		}

		rp := image.NewPaletted(rrec, scUsedP)
		draw.Draw(rp, rrec, rImage, image.ZP, draw.Src)
		outGif.Image = append(outGif.Image, rp)
		outGif.Delay = append(outGif.Delay, g.Img.Delay[u])
		u++
	}

	outGif.Config.Width = int(width)
	outGif.Config.Height = int(height)

	out, err := os.Create(o.OutputFile)
	if err != nil {
		return err
	}
	defer out.Close()

	err = gif.EncodeAll(out, &outGif)
	if err != nil {
		return err
	}
	if e != 0 {
		return fmt.Errorf("Failed to save %d frames", e)
	}
	return nil
}
