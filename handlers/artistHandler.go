package handlers

import (
	"fmt"
	"musicHub/utils"
	"net/http"
	"sync"
	"time"
)

func Artist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := r.FormValue("id")
	if !utils.IsId(id) {
		http.Error(w, "not found", http.StatusNotFound)
		return

	}
	var (
		artist utils.Artist
		wg     sync.WaitGroup
		mu     sync.Mutex
	)
	errChan := make(chan error, len(utils.Url))
	for index, url := range utils.Url {
		wg.Add(1)
		switch index {
		case "artists":
			go utils.Fetch(fmt.Sprintf("%s/%s", url, id), &artist, &wg, &mu, errChan)
		case "dates":
			go utils.Fetch(fmt.Sprintf("%s/%s", url, id), &artist.Dates, &wg, &mu, errChan)
		case "locations":
			go utils.Fetch(fmt.Sprintf("%s/%s", url, id), &artist.Locations, &wg, &mu, errChan)
		case "relation":
			go utils.Fetch(fmt.Sprintf("%s/%s", url, id), &artist.Relations, &wg, &mu, errChan)
		}

	}

	go func() {
		wg.Wait()
		close(errChan)

	}()
	select {
	case err := <-errChan:
		if err != nil {
			fmt.Println(err)
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
	case <-time.After(10 * time.Second):
		http.Error(w, "request timeout", http.StatusRequestTimeout)
		return
	}

	utils.Templates.ExecuteTemplate(w, "artist.html", artist)
}
