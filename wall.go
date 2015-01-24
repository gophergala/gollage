package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"image"
	"net/http"
	"strings"
)

const GridWidth = 960

type Wall struct {
	Images  []Image
	Url     string
	Name    string
	Heights []int
}

type Image struct {
	Pic     image.Image
	XOffset int
	YOffset int
}

func (w *Wall) AddImage(pic image.Image) {

	image := Image{pic, 0, 0}
	w.Images = append(w.Images)

}

func newWallHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	name = strings.Replace(name, " ", "_", -1)

	if len(name) == 0 {
		fmt.Println("No name entered")
		// Let them know they dun goofed
		http.Redirect(w, r, "/error", 302)
		return
	}
	// If it exists, they can't have it
	if _, ok := walls["foo"]; ok {
		// Sorry brah, this wall's taken
	} else {
		err := NewWallBucket(name)
		if err != nil {
			fmt.Println("Error making bucket", err)
			// Let them know we couldn't persist it
			http.Redirect(w, r, "/error", 302)
			return
		} else {
			// Don't make the wall until we're sure we can persist it
			walls[name] = Wall{
				Images: []Image{},
				Name:   name,
			}
			fmt.Println("Made new wall:", name)
			http.Redirect(w, r, "/wall/"+name, 302)
			return
		}
	}
}

func wallHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Make sure the wall exists
	wall, ok := walls[vars["id"]]
	if ok {
		data := struct {
			Wall Wall
		}{
			wall,
		}
		err := templates.ExecuteTemplate(w, "wall.html", data)
		if err != nil {
			fmt.Println("Error executing template:", err)
		}
	} else {
		fmt.Println("Tried to view non-existent wall")
		http.Redirect(w, r, "/error", 302)
	}
}
