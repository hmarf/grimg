package rimg

import (
	"fmt"
	"image"
	"image/gif"
	"os"
)

func Grimg() {
	// compression rate
	comp := 0.5

	filepath := "./img/test.gif"
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		fmt.Printf("\x1b[31m%s\x1b[0m", "can not open file")
		return
	}

	img, format, err := image.Decode(file)
	if err != nil {
		fmt.Printf("\x1b[31m%s\x1b[0m", "File not in image format")
		return
	}

	r := img.Bounds()
	switch format {
	case "png":
		fmt.Println("png")
		p := pngService{Img: &img}
		err = p.resize(uint(float64(r.Dx())*comp), uint(float64(r.Dy())*comp))
		if err != nil {
			fmt.Printf("\x1b[31m%s\x1b[0m", err)
			return
		}
	case "jpeg":
		fmt.Println("jpeg")
		j := jpegService{Img: &img}
		err = j.resize(uint(float64(r.Dx())*comp), uint(float64(r.Dy())*comp))
		if err != nil {
			fmt.Printf("\x1b[31m%s\x1b[0m", err)
			return
		}
	case "gif":
		file.Seek(0, 0)
		gifimg, err := gif.DecodeAll(file)
		if err != nil {
			fmt.Printf("\x1b[31m%s\x1b[0m", err)
			return
		}
		fmt.Println("gif")
		g := gifService{Img: gifimg}
		err = g.resize(uint(float64(r.Dx())*comp), uint(float64(r.Dy())*comp), comp)
		if err != nil {
			fmt.Printf("\x1b[31m%s\x1b[0m", err)
			return
		}
	default:
		fmt.Printf("\x1b[31m%s\x1b[0m", "The format is not supported")
		return
	}
}
