package images

import (
	"image"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func ReadImageFile(path string) (image.Image, error) {
	return ReadImageFileAndResize(path, 0, 0)
}

func ReadImageFileAndResize(path string, width, height uint) (image.Image, error) {
	img, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer img.Close()
	original, err := png.Decode(img)
	if err != nil {
		return nil, err
	}
	if width == 0 && height == 0 {
		return original, err
	}

	result := resize.Resize(width, height, original, resize.Lanczos3)
	return result, nil
}
