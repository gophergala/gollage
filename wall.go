package main

import (
	"fmt"
	"image"
	"net/http"
	"strings"
)

type Wall struct {
	Images []image.Image
	Url    string
	Name   string
}

func (w *Wall) AddImage(pic image.Image) {

}

func newWallHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	name = strings.Replace(name, " ", "_", -1)
	// If it exists, they can't have it
	if _, ok := walls["foo"]; ok {
		// Sorry brah, this wall's taken
	} else {
		err := NewWallBucket(name)
		if err != nil {
			fmt.Println("Error making bucket", err)
			// Let them know we couldn't persist it
			http.Redirect(w, r, "/error", 500)
		} else {
			// Don't make the wall until we're sure we can persist it
			walls[name] = Wall{
				Images: []image.Image{},
				Name:   name,
			}
			fmt.Println("Made new wall", name)
			http.Redirect(w, r, "/wall/"+name, 301)
		}
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
