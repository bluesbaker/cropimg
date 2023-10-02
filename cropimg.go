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

type ProgramFlags struct {
	Source string
	Output string
	Width  int
	Height int
	Left   int
	Top    int
}

var flags ProgramFlags = ProgramFlags{}

func init() {
	// init default flags from ./flags.go
	initFlags(&flags)
}

func main() {
	flag.Parse()

	// check image(s) source
	if flags.Source == "" {
		fmt.Printf("Image source '%s' is empty\n", flags.Source)
		os.Exit(1)
	}

	// get image(s)
	images, err := filepath.Glob(flags.Source)
	if err != nil {
		fmt.Println("Pattern is not readable", flags.Source)
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
		imageFile = imageutil.Crop(imageFile, flags.Width, flags.Height, flags.Left, flags.Top)

		// save
		filePath := formatFilePath(imagePath, flags.Output, i+1)
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
