package rimg

import (
	"fmt"
	"image"
	"image/gif"
	_ "image/gif"
	"log"
	"os"
)

type Option struct {
	InputFile   string
	OutputFile  string
	Compression float64 // Compression ratio
}

func Grimg(o Option) {

	file, err := os.Open(o.InputFile)
	defer file.Close()
	if err != nil {
		log.Fatalf("\x1b[31m%s\x1b[0m", err)
	}

	img, format, err := image.Decode(file)
	if err != nil {
		log.Fatalf("\x1b[31m%s\x1b[0m", err)
	}

	r := img.Bounds()
	switch format {
	case "png":
		if o.OutputFile == "output" {
			o.OutputFile = "output.png"
		}
		p := pngService{Img: &img}
		err = p.resize(uint(float64(r.Dx())*o.Compression), uint(float64(r.Dy())*o.Compression), o)
		if err != nil {
			log.Fatalf("\x1b[31m%s\x1b[0m", err)
		}
	case "jpeg":
		if o.OutputFile == "output" {
			o.OutputFile = "output."
		}
		j := jpegService{Img: &img}
		err = j.resize(uint(float64(r.Dx())*o.Compression), uint(float64(r.Dy())*o.Compression), o)
		if err != nil {
			log.Fatalf("\x1b[31m%s\x1b[0m", err)
		}
	case "gif":
		if o.OutputFile == "output" {
			o.OutputFile = "output.gif"
		}
		file.Seek(0, 0)
		gifimg, err := gif.DecodeAll(file)
		if err != nil {
			log.Fatalf("\x1b[31m%s\x1b[0m", err)
		}
		fmt.Println("gif")
		g := gifService{Img: gifimg}
		fmt.Println(r)
		err = g.resize(uint(float64(r.Dx())*o.Compression), uint(float64(r.Dy())*o.Compression), o)
		if err != nil {
			log.Fatalf("\x1b[31m%s\x1b[0m", err)
		}
	default:
		fmt.Printf("\x1b[31m%s\x1b[0m", "The format is not supported")
		return
	}
}
