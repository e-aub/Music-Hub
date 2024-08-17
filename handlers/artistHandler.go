package handlers

import (
	"fmt"
	"musicHub/utils"
	"net/http"
	"sync"
	"time"
)

func Artist(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Println(r.URL.Path)
	id := r.FormValue("id")
	fmt.Println(id)

	var artist utils.Artist
	var wg sync.WaitGroup
	var mu sync.Mutex
	for index, url := range utils.Url {
		wg.Add(1)
		switch index {
		case "artists":
			go utils.Fetch(fmt.Sprintf("%s/%s", url, id), &artist, &wg, &mu)
		case "dates":
			go utils.Fetch(fmt.Sprintf("%s/%s", url, id), &artist.Dates, &wg, &mu)
		case "locations":
			go utils.Fetch(fmt.Sprintf("%s/%s", url, id), &artist.Locations, &wg, &mu)
		case "relation":
			go utils.Fetch(fmt.Sprintf("%s/%s", url, id), &artist.Relations, &wg, &mu)
		}
	}
	wg.Wait()
	fmt.Println(artist)
	fmt.Println(time.Since(start))
}
