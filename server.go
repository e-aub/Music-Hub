package main

import (
	"fmt"
	"html/template"
	"log"
	"musicHub/handlers"
	"musicHub/utils"
	"net/http"
)

func main() {
	var err error
	utils.Templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalln(err)
	}
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.Artists)
	http.HandleFunc("/artist", handlers.Artist)
	fmt.Println("\033[32;1mStarting server at port :8080\033[0m")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
