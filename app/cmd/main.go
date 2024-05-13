package main

import (
	"app/config"
	"app/internal"
	"fmt"
)

func main() {
	fmt.Println(internal.LoadImages(config.PathToBody))
}
