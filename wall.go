package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/png"
	"net/http"
	"strings"

	_ "code.google.com/p/vp8-go/webp"
	_ "image/jpeg"
)

const GridWidth = 960
const GridHeight = 540
const HorizontalMargin = 5
const VerticalMargin = 5

type Wall struct {
	Images  []*Image
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
	w.ClearPositioning()
	w.Images = append(w.Images, &Image{Pic: pic})
	w.Run(205)
	w.DrawWall()
}

func newWallHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	name = strings.Replace(name, " ", "", -1)

	if len(name) == 0 {
		fmt.Println("No name entered")
		// Let them know they dun goofed
		http.Redirect(w, r, "/error", 302)
		return
	}
	// If it exists, they can't have it
	if _, ok := walls[name]; ok {
		// Sorry brah, this wall's taken
		http.Redirect(w, r, "/error", 302)
		return
	} else {
		//
		err := NewWallBucket(name)
		if err != nil {
			fmt.Println("Error making bucket", err)
			// Let them know we couldn't persist it
			http.Redirect(w, r, "/error", 302)
			return
		} else {
			//Don't make the wall until we're sure we can persist it
			walls[name] = &Wall{
				Images: []*Image{},
				Name:   name,
			}
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
			Wall   Wall
			Width  int
			Height int
			Host   string
		}{
			*wall,
			GridWidth,
			GridHeight,
			r.Host,
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

func GetHeight(images []*Image, width int) int {
	width -= len(images) * HorizontalMargin
	height := 0
	for _, img := range images {
		height += img.Pic.Bounds().Dx() / img.Pic.Bounds().Dy()
	}
	return width / height
}

func (w *Wall) CalcYOffset(rowNum int) int {
	yOffset := 0
	if rowNum > 0 {
		for _, rowHeight := range w.Heights[:rowNum] {
			yOffset += rowHeight
		}
	}
	return yOffset
}

func (w *Wall) SetRow(images []*Image, height, rowNum int) {
	w.Heights = append(w.Heights, height+VerticalMargin)
	xOffset, yOffset := 0, w.CalcYOffset(rowNum)
	for _, image := range images {
		bounds := image.Pic.Bounds()
		image.DispWidth = height * bounds.Dx() / bounds.Dy()
		image.DispHeight = height

		image.XOffset = xOffset
		image.YOffset = yOffset

		xOffset += image.DispWidth + HorizontalMargin
	}
}

func (w *Wall) Run(maxHeight int) {
	var slice []*Image
	var height int
	n := 0
	images := w.Images
OuterLoop:
	for len(images) > 0 {
		for i := 1; i < len(images)+1; i++ {
			slice = images[:i]
			height = GetHeight(slice, GridWidth)
			if height < maxHeight {
				w.SetRow(slice, height, n)
				images = images[i:]
				n++
				continue OuterLoop
			}
		}
		w.SetRow(images, Min(maxHeight, height), n)
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

func (w *Wall) DrawWall() {

	b := image.Rect(0, 0, GridWidth, GridHeight)
	m := image.NewRGBA(b)
	var original = make(chan Image, 100)
	var resized = make(chan Image, 100)

	// We make worker threads for resizing images whenever we draw a new wall
	for i := 0; i < 5; i++ {
		go ResizeWorker(original, resized)
	}

	for _, img := range w.Images {
		original <- *img
	}
	close(original)

	for _, _ = range w.Images {
		img := <-resized
		loc := image.Rect(img.XOffset, img.YOffset, img.XOffset+img.DispWidth, img.YOffset+img.DispHeight)
		draw.Draw(m, loc, img.Pic, image.ZP, draw.Src)
	}

	out := new(bytes.Buffer)
	encoder := png.Encoder{png.BestCompression}
	encoder.Encode(out, m)
	AddWallImage(w.Name, out)
}

func (w *Wall) ClearPositioning() {
	w.Heights = []int{}
	for _, img := range w.Images {
		img.XOffset = 0
		img.YOffset = 0
		img.DispWidth = 0
		img.DispHeight = 0
	}
}

func ResizeWorker(originals <-chan Image, resized chan<- Image) {
	for img := range originals {
		newImage := Image{
			XOffset:    img.XOffset,
			YOffset:    img.YOffset,
			DispWidth:  img.DispWidth,
			DispHeight: img.DispHeight,
		}
		newImage.Pic = resize.Resize(uint(img.DispWidth), uint(img.DispHeight), img.Pic, resize.NearestNeighbor)
		resized <- newImage
	}
}
