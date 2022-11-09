package main

import (
	"fmt"
	"groupie-tracker/controller"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", controller.MainPage)
	http.HandleFunc("/artist/", controller.Get)
	log.Fatal(http.ListenAndServe(":8081", nil))
	fmt.Println("Server running on http://localhost:8081")
}
