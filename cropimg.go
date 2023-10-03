package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	iu "github.com/bluesbaker/cropimg/pkg/imageutil"
)

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

	// get ignored image(s)
	ignoredImages, err := filepath.Glob(flags.Ignore)
	if err != nil {
		fmt.Println("Ignored pattern is not readable", flags.Ignore)
		os.Exit(1)
	}

	// drop ignored image(s)
	for _, ignoredImage := range ignoredImages {
		for i, image := range images {
			if image == ignoredImage {
				images = append(images[:i], images[i+1:]...)
				break
			}
		}
	}

	dirIndexes := make(map[string]int)

	// crop and save image(s)
	for i, img := range images {
		fileInfo := iu.GetImageInfo(img)

		dirIndexes[fileInfo.Dir]++

		// open
		imageFile, imageExt, err := iu.Open(img)
		if err != nil {
			fmt.Println("Open image error:", err)
			os.Exit(1)
		}

		// crop
		imageFile = iu.Crop(imageFile, flags.Width, flags.Height, flags.Left, flags.Top)

		// save
		formatedOutput := iu.FormatedOutput(fileInfo, flags.Output, i+1, dirIndexes[fileInfo.Dir])
		imagePath, err := iu.Save(imageFile, imageExt, formatedOutput)
		if err != nil {
			fmt.Println("Save image error:", err)
			os.Exit(1)
		}
		fmt.Println(imagePath)
	}
}
