package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type RequestBody struct {
	Style string `json:"style"`
	Text  string `json:"text"`
}

func main() {
	port := "8080"
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", loadMainPage)
	mux.HandleFunc("/ascii-art", handlePostAscii)

	fmt.Println(`Succesfully started server at localhost:` + port)
	http.ListenAndServe(":"+port, mux)
}

func loadMainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		w.Write([]byte("Error 404. Page not found."))
		return
	}

	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	tmpl.Execute(w, nil)
}

func handlePostAscii(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	json.Unmarshal(body, &requestBody)
	style := requestBody.Style
	text := requestBody.Text

	file := GenerateASCII(text, style)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(file))
}
