package handlers

// func Artists(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.Error(w, "404 Page not found", http.StatusNotFound)
// 		return
// 	}
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// 	var artists []utils.HomeArtist
// 	err := utils.Fetch("https://groupietrackers.herokuapp.com/api/artists", &artists)
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Println(artists)
// 	err = utils.Templates.ExecuteTemplate(w, "index.html", artists)
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 	}
// }
