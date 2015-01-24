package main

import (
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"math"
	"net/http"
)

type Wall struct {
	Images []image.Image
	Url    string
}

func (w *Wall) AddImage(pic image.Image) {

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

func newWallHandler(w http.ResponseWriter, r *http.Request) {
	status := 200
	// If it exists, they can't have it
	if _, ok := walls["foo"]; ok {
		// Wholly unacceptable
		status = 406
	}
	data := struct {
		Status int
	}{
		status,
	}
	err := templates.ExecuteTemplate(w, "wall.html", data)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}

func wallHandler(w http.ResponseWriter, r *http.Request) {
	status := 200
	// Make sure the wall exists
	wall, ok := walls["foo"]
	if ok {
		// Wholly unacceptable
		status = 406
	}
	data := struct {
		Status int
		Wall   Wall
	}{
		status,
		wall,
	}
	err := templates.ExecuteTemplate(w, "wall.html", data)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}
