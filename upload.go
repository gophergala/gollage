package main

import (
	"fmt"
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
	}

	// the FormFile function takes in the POST input id file
	file, header, err := r.FormFile("file")

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	AddImageToBucket(wall, wallName, header.Filename, file, r.ContentLength)

	fmt.Fprintf(w, "File uploaded successfully : ")
	fmt.Fprintf(w, header.Filename)

	resp.WriteOut(w)
}
