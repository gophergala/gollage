package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"image"
	"net/http"
	"strings"
)

const GridWidth = 960
const GridHeight = 540

type Wall struct {
	Images  []Image
	Url     string
	Name    string
	Heights []int
}

type Image struct {
	Pic        image.Image
	XOffset    int
	YOffset    int
	DispWidth  int
	DispHeight int
}

func (w *Wall) AddImage(pic image.Image) {
	w.Images = append(w.Images, Image{pic, 0, 0, 0, 0})
	w.Run(205)
	fmt.Printf("%+v\n", w)
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

// Stole this javascript from http://blog.vjeux.com/wp-content/uploads/2012/05/google-layout.html

func GetHeight(images []Image, width int) int {
	width -= len(images) * 5
	height := 0
	for _, image := range images {
		height += image.Pic.Bounds().Dx() / image.Pic.Bounds().Dy()
	}
	return width / height
}

func (w *Wall) SetHeight(images []Image, height int) {
	w.Heights = append(w.Heights, height)
	for _, image := range images {
		bounds := image.Pic.Bounds()
		image.DispWidth = height * bounds.Dx() / bounds.Dy()
		image.DispHeight = height
	}
}

func (w *Wall) Resize(images []Image, width int) {
	w.SetHeight(images, GetHeight(images, width))
}

func (w *Wall) Run(maxHeight int) {
	size := GridWidth - 50
	n := 0

	images := []Image{}
	copy(images, w.Images)
	var slice []Image
	var height int
OuterLoop:
	for len(images) > 0 {
		for i := 1; i < len(images)+1; i++ {
			slice = images[0:i]
			height = GetHeight(slice, size)
			if height < maxHeight {
				w.SetHeight(slice, height)
				n++
				images = images[:i]
				continue OuterLoop
			}
		}
		w.SetHeight(slice, Min(maxHeight, height))
		n++
		break
	}
}

func Min(inputs ...int) int {
	smallest := inputs[0]
	for _, num := range inputs {
		if num < smallest {
			smallest = num
		}
	}
	return smallest
}
