package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type SubImage interface {
	SubImage(r image.Rectangle) image.Image
}

func main() {
	// flags
	imageSource := flag.String("source", "", "image(s) source")
	postfix := flag.String("postfix", "_cropped", "cropped image postfix")
	width := flag.Int("width", 0, "width")
	height := flag.Int("height", 0, "height")
	leftOffset := flag.Int("left", 0, "left offset")
	topOffset := flag.Int("top", 0, "top offset")
	flag.Parse()

	// check image source
	if *imageSource == "" {
		fmt.Printf("Image source '%s' is empty\n", *imageSource)
		os.Exit(1)
	}

	images, err := filepath.Glob(*imageSource)
	if err != nil {
		fmt.Println("Pattern is not readable", *imageSource)
		os.Exit(1)
	}

	for _, img := range images {
		cropImage(img, *width, *height, *leftOffset, *topOffset, *postfix)
	}
}

func cropImage(filePath string, width, height, left, top int, postfix string) {
	var originalImage image.Image
	ext := filepath.Ext(filePath)
	fileName := strings.Split(filepath.Base(filePath), ext)[0]

	originalImageFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer originalImageFile.Close()

	switch {
	case ext == ".png":
		originalImage, err = png.Decode(originalImageFile)
		if err != nil {
			panic(err)
		}
	case (ext == ".jpeg") || (ext == ".jpg"):
		originalImage, err = jpeg.Decode(originalImageFile)
		if err != nil {
			panic(err)
		}
	default:
		fmt.Printf("file extension is bad!")
		os.Exit(1)
	}

	cropSize := image.Rect(0, 0, width, height)
	cropSize = cropSize.Add(image.Point{left, top})
	croppedImage := originalImage.(SubImage).SubImage(cropSize)

	croppedImageFile, err := os.Create(filepath.Join(filepath.Dir(filePath), fmt.Sprintf("%s%s%s", fileName, postfix, ext)))
	if err != nil {
		panic(err)
	}
	defer croppedImageFile.Close()

	switch {
	case ext == ".png":
		if err := png.Encode(croppedImageFile, croppedImage); err != nil {
			panic(err)
		}
	case (ext == ".jpeg") || (ext == ".jpg"):
		if err := jpeg.Encode(croppedImageFile, croppedImage, nil); err != nil {
			panic(err)
		}
	}
}
