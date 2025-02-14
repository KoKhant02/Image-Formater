package resize

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

var FormatSizes = map[string][2]float64{
	"license":     {3.5, 4.5},
	"driver":      {5.0, 5.0},
	"national_id": {2.5, 3.0},
}

type ImageResizer interface {
	Resize(inputPath, formatType string, dpi uint) error
}

type Resizer struct{}

func NewResizer() *Resizer {
	return &Resizer{}
}

func cmToPixels(cm float64, dpi uint) uint {
	return uint((cm / 2.54) * float64(dpi))
}

func (r *Resizer) Resize(inputPath, formatType string, dpi uint) error {
	size, exists := FormatSizes[formatType]
	if !exists {
		return fmt.Errorf("invalid format type: %s", formatType)
	}

	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	newWidth := cmToPixels(size[0], dpi)
	newHeight := cmToPixels(size[1], dpi)

	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

	outputPath := fmt.Sprintf("%s_%s_%d.%s", strings.TrimSuffix(inputPath, filepath.Ext(inputPath)), formatType, dpi, format)

	if err := saveImage(outputPath, resizedImg, format); err != nil {
		return err
	}

	fmt.Printf("Resized image saved: %s\n", outputPath)
	return nil
}

func saveImage(outputPath string, img image.Image, format string) error {
	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	switch format {
	case "jpeg", "jpg":
		return jpeg.Encode(outFile, img, &jpeg.Options{Quality: 100})
	case "png":
		return png.Encode(outFile, img)
	default:
		return errors.New("unsupported image format")
	}
}

type App struct {
	resizer ImageResizer
}

func NewApp(resizer ImageResizer) *App {
	return &App{resizer: resizer}
}

func (a *App) Run(imgPath, formatType string, dpi uint) {
	if err := a.resizer.Resize(imgPath, formatType, dpi); err != nil {
		fmt.Println("Error:", err)
	}
}
