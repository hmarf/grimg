package rimg

import (
	"fmt"
	"image"
	"image/gif"
	"log"
	"os"
)

func Grimg() {
	// compression rate
	comp := 0.5

	filepath := "./img/test.gif"
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Printf("%v", "ファイルを開くことができません")
		return
	}

	img, format, err := image.Decode(file)
	if err != nil {
		log.Printf("%v", "画像形式ではありません")
		return
	}

	r := img.Bounds()
	switch format {
	case "png":
		fmt.Println("png")
		p := pngService{Img: &img}
		p.resize(uint(float64(r.Dx())*comp), uint(float64(r.Dy())*comp))
	case "jpeg":
		fmt.Println("jpeg")
		j := jpegService{Img: &img}
		j.resize(uint(float64(r.Dx())*comp), uint(float64(r.Dy())*comp))
	case "gif":
		file.Seek(0, 0)
		gifimg, err := gif.DecodeAll(file)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		fmt.Println("gif")
		g := gifService{Img: gifimg}
		g.resize(uint(float64(r.Dx())*comp), uint(float64(r.Dy())*comp), comp)
	default:
		log.Printf("%v", "対応していないフォーマットです")
	}
}
