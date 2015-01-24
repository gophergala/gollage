package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const ImageSize = 300

func main() {
	go h.run()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/ws", serveWs)
	r.HandleFunc("/wall/{id}", uploadHandler)

	http.Handle("/", r)

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
