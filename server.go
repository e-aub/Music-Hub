package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"text/template"
)

type (
	Artist struct {
		Id                int      `json:"id"`
		Image             string   `json:"image"`
		Name              string   `json:"Name"`
		Members           []string `json:"members"`
		CreationDate      int      `json:"creationDate"`
		FirstAlbum        string   `json:"firstAlbum"`
		LocationsAndDates map[string][]string
	}

	Relation struct {
		LocationAndDates map[string][]string `json:"datesLocations"`
	}
)

var (
	indexRegex = regexp.MustCompile(`^\{\"index\":(.*)}`)
	Artists    []Artist
	Data       bytes.Buffer
	Relations  []Relation
	Templates  *template.Template
)

func serve(w http.ResponseWriter, r *http.Request) {
	Templates.ExecuteTemplate(w, "index.html", Artists)
}

func fetcher(url string, typ any) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	Data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if match := indexRegex.FindAllSubmatch(Data, -1); match != nil {
		Data = match[0][1]
	}
	err = json.Unmarshal(Data, typ)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := fetcher("https://groupietrackers.herokuapp.com/api/artists", &Artists)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fetcher("https://groupietrackers.herokuapp.com/api/relation", &Relations)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(Relations) == 52 && len(Artists) == 52 {
		for i := 0; i < 52; i++ {
			Artists[i].LocationsAndDates = Relations[i].LocationAndDates
		}
	} else {
		fmt.Println("Invalid data")
	}

	Templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalln(err)
	}
	http.HandleFunc("/", serve)
	http.ListenAndServe(":8080", nil)
}
