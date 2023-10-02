package imageutil

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

type SubImage interface {
	SubImage(r image.Rectangle) image.Image
}

func Open(filePath string) (image.Image, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	return image.Decode(file)
}

func Save(imageFile image.Image, encodeExt, filePath string) (string, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	switch encodeExt {
	case "jpg":
	case "jpeg":
		err = jpeg.Encode(file, imageFile, nil)
	case "png":
		err = png.Encode(file, imageFile)
	case "gif":
		err = gif.Encode(file, imageFile, nil)
	}

	if err != nil {
		return "", err
	}
	return filePath, nil
}

func Crop(imageFile image.Image, width, height, left, top int) image.Image {
	cropSize := image.Rect(0, 0, width, height)
	cropSize = cropSize.Add(image.Point{left, top})

	return imageFile.(SubImage).SubImage(cropSize)
}
