package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"musicHub/utils"
	"net/http"
	"os"
)

func Artists(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	var artists []utils.HomeArtist
	response, err := http.Get(utils.Url["artists"])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	Data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	err = json.Unmarshal(Data, &artists)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	err = utils.Templates.ExecuteTemplate(w, "index.html", artists)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
