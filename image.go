// Stole the majority of this from http://blog.golang.org/go-image-package
package main

import (
	"image"
	"image/png"
	"io"

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
