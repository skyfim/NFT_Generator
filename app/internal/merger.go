package internal

import (
	"app/config"
	"image"
	"image/draw"
	"log"
)

func MergeLayers() {
	asset, err := LoadImage(config.PathToAsset)
	if err != nil {
		log.SetPrefix("merger.go: ")
		log.Fatal(err)
	}
	dummy, err := LoadImage(config.PathToDummy)
	if err != nil {
		log.SetPrefix("merger.go: ")
		log.Fatal(err)
	}

	r := image.Rectangle{}
	newImage := image.NewRGBA()
	draw.Draw()
}
