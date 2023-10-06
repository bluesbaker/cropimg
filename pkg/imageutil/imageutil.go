package imageutil

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ImageInfo struct {
	Name string
	Ext  string
	Dir  string
}

type SubImage interface {
	SubImage(r image.Rectangle) image.Image
}

func Open(filePath string) (image.Image, *ImageInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, nil, err
	}

	baseName := strings.Split(filepath.Base(filePath), ".")
	imgInfo := &ImageInfo{
		Name: baseName[0],
		Ext:  baseName[1],
		Dir:  filepath.Dir(filePath),
	}

	return img, imgInfo, nil
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

func (i *ImageInfo) Format(format string, index, localIndex int) string {
	replaceMap := map[string]string{
		"{dir}":   "%[1]s",
		"{name}":  "%[2]s",
		"{ext}":   "%[3]s",
		"{time}":  "%[4]s",
		"{date}":  "%[5]s",
		"{index}": "%[6]d",
		"{local}": "%[7]d",
	}
	output := format
	now := time.Now()
	timeNow := now.Format("15-04-05")
	dateNow := now.Format("02.01.2006")

	for key, value := range replaceMap {
		output = strings.ReplaceAll(output, key, value)
	}

	return fmt.Sprintf(output, i.Dir, i.Name, i.Ext, timeNow, dateNow, index, localIndex)
}
