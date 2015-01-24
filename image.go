// Stole the majority of this from http://blog.golang.org/go-image-package
package main

import (
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/png"
	"io"
	"math"

	_ "code.google.com/p/vp8-go/webp"
	_ "image/jpeg"
)

// convertToPNG converts from any recognized format to PNG.
func convertToPNG(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return png.Encode(w, img)
}

func Resize(totalPix int, pic *image.Image) error {
	bounds := (*pic).Bounds()
	// Do actual image manipulations (with ImageMagick?)
	if bounds.Dx() == 0 || bounds.Dy() == 0 {
		return errors.New("One or more of your dimensions is zero")
	}
	// Ratio
	ratio := bounds.Dx() / bounds.Dy()
	width := uint(math.Floor(math.Sqrt(float64(ratio * totalPix))))
	newPic := resize.Resize(width, 0, *pic, resize.Lanczos3)
	pic = &newPic
	return nil
}
