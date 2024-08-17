package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

type (
	Artist struct {
		Id           int       `json:"id"`
		Image        string    `json:"image"`
		Name         string    `json:"Name"`
		Members      []string  `json:"members"`
		CreationDate int       `json:"creationDate"`
		FirstAlbum   string    `json:"firstAlbum"`
		Locations    Locations `json:"-"`
		Dates        Dates     `json:"-"`
		Relations    Relation  `json:"-"`
	}

	HomeArtist struct {
		Id    int    `json:"id"`
		Image string `json:"image"`
		Name  string `json:"Name"`
	}

	Relation struct {
		LocationAndDates map[string][]string `json:"datesLocations"`
	}

	Locations struct {
		Locations []string `json:"locations"`
	}

	Dates struct {
		Dates []string `json:"dates"`
	}
)

var (
	indexRegex = regexp.MustCompile(`^\{\"index\":(.*)}`)
	Data       bytes.Buffer
	Relations  []Relation
	Templates  *template.Template
	Url        = map[string]string{
		"artists":   "https://groupietrackers.herokuapp.com/api/artists",
		"locations": "https://groupietrackers.herokuapp.com/api/locations",
		"dates":     "https://groupietrackers.herokuapp.com/api/dates",
		"relation":  "https://groupietrackers.herokuapp.com/api/relation",
	}
)

func Fetch(url string, typ any, wg *sync.WaitGroup, mu *sync.Mutex) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	Data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	if match := indexRegex.FindAllSubmatch(Data, -1); match != nil {
		Data = match[0][1]
	}
	mu.Lock()
	err = json.Unmarshal(Data, typ)
	mu.Unlock()
	if err != nil {
		fmt.Println(err)
	}
	defer wg.Done()
}
