package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Stole this from https://www.socketloop.com/tutorials/golang-upload-file
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	wallName := vars["id"]
	wall, ok := walls[wallName]
	// You can't add an image to a wall that doesn't exist
	if ok {
		// Wall exists
	} else {
		// Wall doesn't exist
		fmt.Println("Uh where are you trying to put this?")
		// Let them know they dun goofed
		http.Redirect(w, r, "/error", 302)
		return
	}

	// the FormFile function takes in the POST input id file
	file, _, err := r.FormFile("file")

	if err != nil {
		fmt.Println("Failed to get file from form: " + err.Error())
		// Let them know they dun goofed
		http.Redirect(w, r, "/error", 302)
		return
	}

	defer file.Close()

	if err != nil {
		fmt.Println("Failed to decode image: " + err.Error())
		// Let them know they dun goofed
		http.Redirect(w, r, "/error", 302)
		return
	}

	// Buf will hold the resized raw image data
	buf := new(bytes.Buffer)
	img, err := Normalize(ImageSize, file, buf)
	if err != nil {
		fmt.Println("Failed to normalize image: " + err.Error())
		// Let them know they dun goofed
		http.Redirect(w, r, "/error", 302)
		return
	}

	wall.AddImage(img)
	http.Redirect(w, r, "/wall/"+wallName, 302)
}
