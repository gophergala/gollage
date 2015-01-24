// Stole the majority of this from http://blog.golang.org/go-image-package
package main

// Process for adding an image to a wall:
// 1. Image is uploaded to server
// 2. Image is converted to PNG and normalized
// 3. Image is uploaded to AWS...for some reason
// 4. Image is added to Wall
// 5. Wall regenerates main image, uploads to AWS

// Process for zooming in:
// 1. Get target zoom area
// 2. Do some math or something
// 3. Generate magical new image
// 4. Something?

import (
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/png"
	"io"
	"math"
)

// convertToPNG converts from any recognized format to PNG.
func ConvertToPNG(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return png.Encode(w, img)
}

func Resize(totalPix int, pic image.Image) (image.Image, error) {
	bounds := pic.Bounds()
	if bounds.Dx() == 0 || bounds.Dy() == 0 {
		return nil, errors.New("One or more of your dimensions is zero")
	}
	// Ratio
	ratio := bounds.Dx() / bounds.Dy()
	width := uint(math.Floor(math.Sqrt(float64(ratio * totalPix))))
	return resize.Resize(width, 0, pic, resize.Lanczos3), nil
}

// Stole this javascript from http://blog.vjeux.com/wp-content/uploads/2012/05/google-layout.html
// Convert this shit to Go
/*
$(function () {


func (w *Wall) GetHeight(images, width) {
  width -= len(images) * 5;
	height := 0;
	for i, image := range w.Images {
    height += image.Pic.Bounds().Dx() / image.Pic.Bounds().Dy()
  }
  return width / height;
}

func (w *Wall) SetHeight(images, height) {
  w.Heights = append(w.Heights, height)
		for i, image := range w.Images {
			height += image.Pic.Bounds().Dx() / image.Pic.Bounds().Dy()
    $(images[i]).css({
      width: height * $(images[i]).data('width') / $(images[i]).data('height'),
      height: height
    });
    $(images[i]).attr('src', $(images[i]).attr('src').replace(/w[0-9]+-h[0-9]+/, 'w' + $(images[i]).width() + '-h' + $(images[i]).height()));
	}
}

function resize(images, width) {
  setheight(images, getheight(images, width));
}

func (w *Wall) Run(maxHeight int) {
	size := GridWidth - 50

  n = 0;

	queue = []Image
	copy(queue, w.Images)
	var slice []Image

	OuterLoop:
  for len(queue) > 0 {
		for i := 1; i < len(queue) + 1; i++ {
			slice := queue[0:i]
			height := GetHeight(slice, size)
      if height < maxHeight {
        SetHeight(slice, height)
        n++
				queue = queue[:i]
				continue OuterLoop
      }
    }
    SetHeight(slice, Min(maxHeight, height))
    n++
    break
  }
  console.log(n);
}

func Min(inputs ...int) int {
	smallest = inputs[0]
	for _, val := range inputs {
		if val < smallest {
			smallest = val
		}
	}
	return smallest
}

window.addEventListener('resize', function () { run(205); });
$(function () { run(205); });

});
*/
