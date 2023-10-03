package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type ProgramFlags struct {
	Source string
	Output string
	Width  int
	Height int
	Left   int
	Top    int
	Ignore string
}

func initFlags(flags *ProgramFlags) {
	const (
		defaultSource = ""
		defaultOutput = "{dir}/{name}_cropped.{ext}"
		defaultWidth  = 0
		defaultHeight = 0
		defaultLeft   = 0
		defaultTop    = 0
		defaultIgnore = ""
	)
	const (
		usageSource = "image(s) source:\n" +
			"\t\t\"./<image>.jpg\" - .jpg-image\n" +
			"\t\t\"./*.png\" - .png-images from current directory\n" +
			"\t\t\"./**/**/*.gif\" - .gif-images from deep(2) directories"
		usageOutput = "output file format:\n" +
			"\t\t{dir} - directory\n" +
			"\t\t{name} - file name\n" +
			"\t\t{ext} - file extension\n" +
			"\t\t{time} - current time(24-59-59)\n" +
			"\t\t{date} - current date(01.02.2003)\n" +
			"\t\t{index} - file index\n" +
			"\t\t{local} - local file index"
		usageWidth  = "width"
		usageHeight = "height"
		usageLeft   = "left offset"
		usageTop    = "top offset"
		usageIgnore = "ignored image(s) source"
	)

	flag.StringVar(&flags.Source, "source", defaultSource, usageSource)
	flag.StringVar(&flags.Source, "s", defaultSource, usageSource)

	flag.StringVar(&flags.Output, "output", defaultOutput, usageOutput)
	flag.StringVar(&flags.Output, "o", defaultOutput, usageOutput)

	flag.IntVar(&flags.Width, "width", defaultWidth, usageWidth)
	flag.IntVar(&flags.Width, "w", defaultWidth, usageWidth)

	flag.IntVar(&flags.Height, "height", defaultHeight, usageHeight)
	flag.IntVar(&flags.Height, "h", defaultHeight, usageHeight)

	flag.IntVar(&flags.Left, "left", defaultLeft, usageLeft)
	flag.IntVar(&flags.Left, "l", defaultLeft, usageLeft)

	flag.IntVar(&flags.Top, "top", defaultTop, usageTop)
	flag.IntVar(&flags.Top, "t", defaultTop, usageTop)

	flag.StringVar(&flags.Ignore, "ignore", defaultIgnore, usageIgnore)
	flag.StringVar(&flags.Ignore, "i", defaultIgnore, usageIgnore)

	flag.Usage = printUsage
}

func printUsage() {
	type usageStruct struct {
		Name      string
		Usage     string
		Value     string
		ValueType string
	}
	usageList := make(map[string]*usageStruct)

	fmt.Printf("Usage of %s\n", os.Args[0])
	fmt.Printf("Example: .\\%s -s \".\\**\\*.png\""+
		" -w 100 -h 200 -l 30 -t 40"+
		" -o \"{dir}\\{name}_{time}.{ext}\" \n",
		filepath.Base(os.Args[0]))
	fmt.Println("Flags:")

	flag.VisitAll(func(f *flag.Flag) {
		valueType, _ := flag.UnquoteUsage(f)
		if u := usageList[f.Usage]; u == nil {
			usageList[f.Usage] = &usageStruct{
				Name:      "-" + f.Name,
				Usage:     f.Usage,
				Value:     f.DefValue,
				ValueType: "<" + valueType + ">",
			}
		} else {
			u.Name = "-" + f.Name + "|" + u.Name
		}
	})

	for _, v := range usageList {
		fmt.Println(v.Name, v.ValueType)
		fmt.Printf("\t%s\n", v.Usage)
		if !(v.Value == "" || v.Value == "0") {
			fmt.Println("\tdefault:", v.Value)
		}
	}

	fmt.Println("\nAuthor: github.com/bluesbaker")
}
