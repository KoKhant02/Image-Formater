package util

import (
	"fmt"
	resize "image-formater/resizer"
)

func ShowFormatDetails() {
	fmt.Println("Available formats and their sizes:")
	for format, size := range resize.FormatSizes {
		fmt.Printf("%s: %.2f cm x %.2f cm\n", format, size[0], size[1])
	}
}
