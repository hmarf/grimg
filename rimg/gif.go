package rimg

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"math"
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

func getGifColor(rec image.Rectangle, img image.Image, w int, h int) color.Palette {

	cUsedM := make(map[color.Color]bool)
	// The color used in the original image
	for x := 1; x <= w; x++ {
		for y := 1; y <= h; y++ {
			if _, ok := cUsedM[img.At(x, y)]; !ok {
				cUsedM[img.At(x, y)] = true
			}
		}
	}
	return keys(cUsedM)

}

func (g *gifService) resizeDisposalNone(width uint, height uint, o Option, outGif *gif.GIF) {

	var pImage image.Image
	oWidth, oheight := g.Img.Image[0].Bounds().Dx(), g.Img.Image[0].Bounds().Dy()
	for i, frame := range g.Img.Image {

		rec := frame.Bounds()

		var rImage image.Image
		if i > 0 {
			// 書き出し用のイメージを準備
			outRect := image.Rectangle{image.Pt(0, 0), pImage.Bounds().Size()}
			out := image.NewRGBA(outRect)

			// 描画する
			// 元画像をまず描く
			dstRect := image.Rectangle{image.Pt(0, 0), pImage.Bounds().Size()}
			draw.Draw(out, dstRect, pImage, image.Pt(0, 0), draw.Src)

			// 上書きする
			srcRect := image.Rectangle{image.Pt(rec.Min.X, rec.Min.Y), image.Pt(rec.Max.X, rec.Max.Y)} //rec.Size()
			draw.Draw(out, srcRect, frame, image.Pt(rec.Min.X, rec.Min.Y), draw.Over)

			pImage = out
			rImage = resize.Resize(
				uint(width),
				uint(height),
				out, resize.Lanczos3)

		} else {
			rImage = resize.Resize(
				uint(math.Floor(float64(rec.Dx())*o.Compression)),
				uint(math.Floor(float64(rec.Dy())*o.Compression)),
				frame, resize.Lanczos3)
			pImage = frame
		}

		rrec := rImage.Bounds()

		c := getGifColor(rrec, pImage, oWidth, oheight)

		rp := image.NewPaletted(rrec, c)
		draw.Draw(rp, rrec, rImage, image.ZP, draw.Src)

		outGif.Image = append(outGif.Image, rp)
		outGif.Delay = append(outGif.Delay, g.Img.Delay[i])
		outGif.Disposal = append(outGif.Disposal, g.Img.Disposal[i])
	}
	return
}

func (g *gifService) resizeDisposalPrevious(width uint, height uint, o Option, outGif *gif.GIF) {

	for i, frame := range g.Img.Image {

		rec := frame.Bounds()

		// rImage := resize.Resize(width, height, frame.SubImage(rec), resize.Lanczos3)
		rImage := resize.Resize(
			uint(math.Floor(float64(rec.Dx())*o.Compression)),
			uint(math.Floor(float64(rec.Dy())*o.Compression)),
			frame.SubImage(rec), resize.Lanczos3)
		rrec := rImage.Bounds()

		c := getGifColor(rrec, frame, rec.Dx(), rec.Dy())

		if i > 0 {
			rrec = image.Rect(
				int(math.Floor(float64(rec.Min.X)*o.Compression)),
				int(math.Floor(float64(rec.Min.Y)*o.Compression)),
				rrec.Dx()+int(math.Floor(float64(rec.Min.X)*o.Compression)),
				rrec.Dy()+int(math.Floor(float64(rec.Min.Y)*o.Compression)))
		}

		rp := image.NewPaletted(rrec, c)
		draw.Draw(rp, rrec, rImage, image.ZP, draw.Src)
		outGif.Image = append(outGif.Image, rp)
		outGif.Delay = append(outGif.Delay, g.Img.Delay[i])
		outGif.Disposal = append(outGif.Disposal, g.Img.Disposal[i])
	}

	return
}

func (g *gifService) resize(width uint, height uint, o Option) error {

	outGif := gif.GIF{}

	switch g.Img.Disposal[0] {
	case gif.DisposalNone:
		g.resizeDisposalNone(width, height, o, &outGif)
	case gif.DisposalPrevious:
		g.resizeDisposalPrevious(width, height, o, &outGif)
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
	return nil
}
