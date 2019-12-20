package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func saveImage(fileName string, img *image.RGBA) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("\x1b[31m%s\x1b[0m\n", "creation of the save destination file failed.")
		return
	}
	defer file.Close()

	if err := jpeg.Encode(file, img, &jpeg.Options{100}); err != nil {
		fmt.Printf("\x1b[31m%s\x1b[0m\n", "Failed to save image.")
		return
	}
}

func resizeImage(img image.Image, width int, height int) image.Image {
	return resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
}

func createRGBAImage(img image.Image) image.Image {
	i := img.Bounds()
	nImage := image.NewRGBA(image.Rect(0, 0, i.Dx(), i.Dy()))
	draw.Draw(nImage, nImage.Bounds(), img, i.Min, draw.Src)
	return nImage
}

func judgeImage(img image.Image, file *os.File, comp float64) *image.RGBA {
	if _, ok := img.(*image.RGBA); ok {
		fmt.Println("RGBA")
		ibounds := img.Bounds()
		return resizeImage(img, int(float64(ibounds.Dx())*comp), int(float64(ibounds.Dy())*comp)).(*image.RGBA)
	} else if _, ok := img.(*image.Paletted); ok {
		fmt.Println("paletted")
		file.Seek(0, 0)
		_, err := gif.DecodeAll(file)
		if err != nil {
			log.Printf("%v", err)
			return nil
		}
		fmt.Println("ok")
		return nil
	} else {
		ibounds := img.Bounds()
		return createRGBAImage(
			resizeImage(img, int(float64(ibounds.Dx())*comp), int(float64(ibounds.Dy())*comp))).(*image.RGBA)
	}
	return nil
}

func main() {
	// compression rate
	comp := 0.1

	filepath := "./image/pokemon.png"
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Printf("%v", "ファイルを開くことができません")
		return
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Printf("%v", "画像形式ではありません")
		return
	}
	judgeImage(img, file, comp)
	// sImage := judgeImage(img, buf, comp)
	// if sImage == nil {
	// 	log.Printf("%v", "対応していないファイル形式です")
	// }
	// saveImage("tes.jpg", sImage)
}
