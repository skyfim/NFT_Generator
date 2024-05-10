package main

import (
	"app/config"
	"app/internal"
	"log"
)

func main() {
	image, err := internal.LoadImage(config.PathToAsset)
	if err != nil {
		log.Fatal(err)
	}
	internal.SaveImage(config.PathToSave, image)
}
