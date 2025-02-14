package main

import (
	"flag"
	"fmt"
	resize "image-formater/resizer"
	"image-formater/util"
)

func main() {
	imgPath := flag.String("img", "", "Path to the input image file")
	formatType := flag.String("format", "", "Format type (license, driver, national_id)")
	dpi := flag.Uint("dpi", 300, "DPI value (300, 600, 1200)")
	showFormats := flag.Bool("formats", false, "Show available format details")

	flag.Parse()

	if *showFormats {
		util.ShowFormatDetails()
		return
	}

	if *imgPath == "" || *formatType == "" {
		fmt.Println("Usage: go run main.go -img <image_path> -format <license|driver|national_id> -dpi <300|600|1200>")
		return
	}

	if *dpi != 300 && *dpi != 600 && *dpi != 1200 {
		fmt.Println("Invalid DPI value. Valid options are: 300, 600, or 1200.")
		return
	}

	resizer := resize.NewResizer()
	app := resize.NewApp(resizer)

	app.Run(*imgPath, *formatType, *dpi)
}
