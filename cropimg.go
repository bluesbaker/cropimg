package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bluesbaker/cropimg/pkg/imageutil"
)

func main() {
	// flags
	imageSource := flag.String("source", "", "image(s) source")
	format := flag.String(
		"format",
		"{dir}/{name}_cropped.{ext}",
		"output file format:\n"+
			"\t{dir} - directory\n"+
			"\t{name} - file name\n"+
			"\t{ext} - file extension\n"+
			"\t{time} - current time(24-59-59)\n"+
			"\t{date} - current date(01.02.2003)\n"+
			"\t{index} - file index\n")
	width := flag.Int("width", 0, "width")
	height := flag.Int("height", 0, "height")
	left := flag.Int("left", 0, "left offset")
	top := flag.Int("top", 0, "top offset")
	flag.Parse()

	// check image(s) source
	if *imageSource == "" {
		fmt.Printf("Image source '%s' is empty\n", *imageSource)
		os.Exit(1)
	}

	// get image(s)
	images, err := filepath.Glob(*imageSource)
	if err != nil {
		fmt.Println("Pattern is not readable", *imageSource)
		os.Exit(1)
	}

	// crop and save image(s)
	for i, imagePath := range images {
		// open
		imageFile, imageExt, err := imageutil.Open(imagePath)
		if err != nil {
			fmt.Println("Open image error:", err)
			os.Exit(1)
		}

		// crop
		imageFile = imageutil.Crop(imageFile, *width, *height, *left, *top)

		// save
		filePath := formatFilePath(imagePath, *format, i+1)
		imagePath, err := imageutil.Save(imageFile, imageExt, filePath)
		if err != nil {
			fmt.Println("Save image error:", err)
			os.Exit(1)
		}
		fmt.Println(imagePath)
	}
}

func formatFilePath(filePath, formatString string, index int) string {
	now := time.Now()
	timeNow := now.Format("15-04-05")
	dateNow := now.Format("02.01.2006")
	baseName := strings.Split(filepath.Base(filePath), ".")
	name := baseName[0]
	ext := baseName[1]
	dir := filepath.Dir(filePath)

	formatedPath := formatString
	formatedPath = strings.ReplaceAll(formatedPath, "{dir}", "%[1]s")
	formatedPath = strings.ReplaceAll(formatedPath, "{name}", "%[2]s")
	formatedPath = strings.ReplaceAll(formatedPath, "{ext}", "%[3]s")
	formatedPath = strings.ReplaceAll(formatedPath, "{time}", "%[4]s")
	formatedPath = strings.ReplaceAll(formatedPath, "{date}", "%[5]s")
	formatedPath = strings.ReplaceAll(formatedPath, "{index}", "%[6]d")
	formatedPath = fmt.Sprintf(formatedPath, dir, name, ext, timeNow, dateNow, index)

	return formatedPath
}
