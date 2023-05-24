package webserver

import "net/http"

func MyWebServer() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "info.json")
}
