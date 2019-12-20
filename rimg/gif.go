package rimg

import (
	"fmt"
	"image"
	"image/gif"
	"log"
	"os"
)

type GifService struct {
	Img *image.Image
}

func (g *GifService) Resize(file *os.File, width uint, height uint) error {
	file.Seek(0, 0)
	_, err := gif.DecodeAll(file)
	if err != nil {
		log.Printf("%v", err)
		return nil
	}

	fmt.Println("ok")
	return nil
}
