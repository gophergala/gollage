package main

import (
	"log"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Host string
	}{
		r.Host,
	}
	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Println("Error executing template:", err)
	}
}
