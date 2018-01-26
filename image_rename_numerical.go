package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type imagenames []string

func captureImages() imagenames {
	images, err := filepath.Glob("*[.jpg][.png]")
	if err != nil {
		log.Fatal(err)
	}
	return images
}

func renameImages(images imagenames) {
	for i, image := range images {
		fmt.Println(image, " -> ", strconv.Itoa(i)+image[len(image)-4:])
		err := os.Rename(image, strconv.Itoa(i)+image[len(image)-4:])
		if err != nil {
			log.Fatal(err)
		}
	}
}
