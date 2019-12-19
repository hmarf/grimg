package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"log"
	"math"
	"os"

	"github.com/nfnt/resize"
)

func saveImage(fileName string, img *image.RGBA64) {
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

func main() {
	// compression rate
	comp := 0.1

	filepath := "./image/test.gif"
	stash, err := os.Open(filepath)
	defer stash.Close()
	if err != nil {
		log.Printf("%v", "ファイルを開くことができません")
		return
	}
	gifImage, err := gif.DecodeAll(stash)
	if err != nil {
		log.Printf("%v", "ファイルは.gif形式ではありません")
		return
	}
	imageConf := gifImage.Config
	width := float64(imageConf.Width)
	height := float64(imageConf.Height)
	fmt.Println(width, height)
	fmt.Println(imageConf)
	for _, frame := range gifImage.Image {
		rec := frame.Bounds()
		_image := frame.SubImage(rec)
		resizedImage := resize.Resize(uint(math.Floor(float64(rec.Dx())*comp)),
			uint(math.Floor(float64(rec.Dy())*comp)),
			_image, resize.Lanczos3)
		saveImage("./image.jpg", resizedImage.(*image.RGBA64))
	}
}
