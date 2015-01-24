package main

import (
	"html/template"
	"log"
	"net/http"
)

var index, _ = template.ParseFiles("templates/index.html")

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := struct {
		Host string
	}{
		r.Host,
	}
	err := index.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
	}
}
