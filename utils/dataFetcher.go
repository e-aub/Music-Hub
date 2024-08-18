package utils

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"net/http"
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
	Data      bytes.Buffer
	Relations []Relation
	Templates *template.Template
	Url       = map[string]string{
		"artists":   "https://groupietrackers.herokuapp.com/api/artists",
		"locations": "https://groupietrackers.herokuapp.com/api/locations",
		"dates":     "https://groupietrackers.herokuapp.com/api/dates",
		"relation":  "https://groupietrackers.herokuapp.com/api/relation",
	}
)

func Fetch(url string, typ any, wg *sync.WaitGroup, mu *sync.Mutex, errchan chan error) {
	response, err := http.Get(url)
	if err != nil {
		errchan <- err
		return
	}
	Data, err := io.ReadAll(response.Body)
	if err != nil {
		errchan <- err
		return
	}
	mu.Lock()
	err = json.Unmarshal(Data, typ)
	mu.Unlock()
	if err != nil {
		errchan <- err
		return
	}
	defer wg.Done()
}

func IsId(id string) bool {
	for _, digit := range id {
		if digit < '0' || digit > '9' {
			return false
		}
	}
	return true
}
