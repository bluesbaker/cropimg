# cropimg
A simple tool for cropping images

## Build
```bash
$ go build -o .\bin\cropimg.exe .\cropimg.go .\flags.go
```

## How to use?
Example:
```bash
$ ...\cropimg.exe -s ".\**\*.png" -w 100 -h 200 -l 30 -t 40 -o "{dir}\{name}_{time}.{ext}"  
```

Flags:
```bash
-source|-s <string>
    image(s) source:
        "./<image>.jpg" - .jpg-image
        "./*.png" - .png-images from current directory
        "./**/**/*.gif" - .gif-images from deep(2) directories

-ignore|-i <string>
    ignored image(s) source:
        "./*_cropped.png" - ignore png-images ending with "_cropped.png"

-width|-w <int>
    width

-height|-h <int>
    height

-top|-t <int>
    top offset

-left|-l <int>
    left offset

-output|-o <string>
    output file format:
        {dir} - directory
        {name} - file name
        {ext} - file extension
        {time} - current time(24-59-59)
        {date} - current date(01.02.2003)
        {index} - file index
        {local} - local file index
    default: {dir}/{name}_cropped.{ext}
```

## Version
0.2.1

### Author
___
Konstantin S.G. <github.com/bluesbaker>