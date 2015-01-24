package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"image"
	"image/png"
	"net/http"

	_ "code.google.com/p/vp8-go/webp"
	_ "image/jpeg"
)

// Stole this from https://www.socketloop.com/tutorials/golang-upload-file
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resp := JSONMessage{
		200, // Start OK
		"Image uploaded successfully",
	}

	wallName := vars["id"]
	wall, ok := walls[wallName]
	// You can't add an image to a wall that doesn't exist
	if ok {
		// Wall exists
	} else {
		// Wall doesn't exist
		resp.Status = 404
		resp.Message = "Uh where are you trying to put this?"
		resp.WriteOut(w)
		return
	}

	// the FormFile function takes in the POST input id file
	file, header, err := r.FormFile("file")

	if err != nil {
		resp.Status = 500
		resp.Message = "Failed to get file from form"
		resp.WriteOut(w)
		return
	}

	img, _, err := image.Decode(file)
	if err != nil {
		resp.Status = 500
		resp.Message = "Failed to decode image"
		resp.WriteOut(w)
		return
	}
	img, err = Resize(ImageSize, img)
	if err != nil {
		resp.Status = 500
		resp.Message = "Failed to resize image"
		resp.WriteOut(w)
		return
	}
	buf := new(bytes.Buffer)
	png.Encode(buf, img)
	AddImageToBucket(wall, wallName, header.Filename, buf, r.ContentLength)

	fmt.Fprintf(w, "File uploaded successfully : ")
	fmt.Fprintf(w, header.Filename)

	resp.WriteOut(w)
}
