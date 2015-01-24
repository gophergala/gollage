package main

import (
	"errors"
	"math"
)

type Wall []Image

type Image struct {
	File   string
	Url    string
	X      int
	Y      int
	Width  int
	Height int
}

func (i *Image) Resize(totalPix int) error {
	// Do actual image manipulations (with ImageMagick?)
	if i.X == 0 || i.Y == 0 {
		return errors.New("One or more of your dimensions is zero")
	}
	ratio := i.X / i.Y
	i.X = int(math.Floor(math.Sqrt(float64(ratio * totalPix))))
	i.Y = int(math.Floor(math.Sqrt(float64(totalPix / ratio))))
	return nil
}
