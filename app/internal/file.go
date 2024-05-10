package internal

import (
	"image"
	"image/png"
	"log"
	"os"
)

func LoadImage(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, err := png.Decode(f)
	return image, err
}

func SaveImage(filepath string, picture image.Image) {
	f, err := os.Create(filepath + `image.png`)
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(f, picture)
}
