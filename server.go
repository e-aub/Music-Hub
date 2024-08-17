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

	// err = utils.Fetch("https://groupietrackers.herokuapp.com/api/relation", &utils.Relations)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if len(utils.Relations) == 52 && len(utils.Artists) == 52 {
	// 	for i := 0; i < 52; i++ {
	// 		utils.Artists[i].LocationsAndDates = utils.Relations[i].LocationAndDates
	// 	}
	// } else {
	// 	fmt.Println("Invalid data")
	// }
	var err error
	utils.Templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalln(err)
	}
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// http.HandleFunc("/", handlers.Artists)
	http.HandleFunc("/artist", handlers.Artist)
	fmt.Println("\033[32;1mStarting server at port :8080\033[0m")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
