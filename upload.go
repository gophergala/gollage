package main

import (
	"bytes"
	"github.com/gorilla/mux"
	"net/http"
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
		resp.Message = "Failed to get file from form: " + err.Error()
		resp.WriteOut(w)
		return
	}

	defer file.Close()

	if err != nil {
		resp.Status = 500
		resp.Message = "Failed to decode image: " + err.Error()
		resp.WriteOut(w)
		return
	}

	// Buf will hold the resized raw image data
	buf := new(bytes.Buffer)
	img, err := Normalize(ImageSize, file, buf)
	if err != nil {
		resp.Status = 500
		resp.Message = "Failed to normalize image: " + err.Error()
		resp.WriteOut(w)
		return
	}

	err = AddImageToBucket(wall, wallName, buf, r.ContentLength)
	if err != nil {
		//resp.Status = 500
		//resp.Message = "Failed to get image onto AWS: " + err.Error()
		//resp.WriteOut(w)
		//return
	}

	wall.AddImage(img)
	resp.Message = "File uploaded successfully: " + header.Filename
	resp.WriteOut(w)
}
