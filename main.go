package main

import (
	"fmt"
	"image"
	"image/draw"
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

func judgeImage(img image.Image, comp float64) *image.RGBA {
	if _, ok := img.(*image.NRGBA); ok {
		fmt.Println("NRGBA")
		ibounds := img.Bounds()
		return createRGBAImage(resizeImage(img, int(float64(ibounds.Dx())*comp), int(float64(ibounds.Dy())*comp))).(*image.RGBA)
	} else if _, ok := img.(*image.RGBA); ok {
		fmt.Println("RGBA")
		ibounds := img.Bounds()
		return resizeImage(img, int(float64(ibounds.Dx())*comp), int(float64(ibounds.Dy())*comp)).(*image.RGBA)
	} else if _, ok := img.(*image.YCbCr); ok {
		fmt.Println("YCbCr")
		ibounds := img.Bounds()
		return createRGBAImage(resizeImage(img, int(float64(ibounds.Dx())*comp), int(float64(ibounds.Dy())*comp))).(*image.RGBA)
	} else if _, ok := img.(*image.Paletted); ok {
		fmt.Println("paletted")
	}
	return nil
}

func main() {
	// compression rate
	comp := 0.1

	filepath := "./image/pokemon.png"
	stash, err := os.Open(filepath)
	defer stash.Close()
	if err != nil {
		log.Printf("%v", "ファイルを開くことができません")
		return
	}
	img, _, err := image.Decode(stash)
	if err != nil {
		log.Printf("%v", "画像形式ではありません")
	}

	sImage := judgeImage(img, comp)
	saveImage("tes.jpg", sImage)

	// gifImage, err := gif.DecodeAll(stash)
	// if err != nil {
	// 	log.Printf("%v", "ファイルは.gif形式ではありません")
	// 	return
	// }
	// imageConf := gifImage.Config
	// width := float64(imageConf.Width)
	// height := float64(imageConf.Height)
	// fmt.Println(width, height)
	// fmt.Println(imageConf)
	// for _, frame := range gifImage.Image {
	// 	rec := frame.Bounds()
	// 	_image := frame.SubImage(rec)
	// 	resizedImage := resize.Resize(uint(math.Floor(float64(rec.Dx())*comp)),
	// 		uint(math.Floor(float64(rec.Dy())*comp)),
	// 		_image, resize.Lanczos3)

	// }
}
