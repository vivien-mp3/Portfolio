package main

import (
	"html/template"
	"log"
	"net/http"
	"bytes"
)

var _temp = template.Must(template.ParseGlob("./templates/*.html"))

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/{$}", HomeHandler)

	fileServer := http.FileServer(http.Dir("./assets"))
	router.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Printf("Serveur: http://localhost:8080")

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatalf("Erreur côté serveur %s", err)
		return
	}

}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	buffer := bytes.Buffer{}

	_temp.ExecuteTemplate(&buffer, name, data)

	w.Header().Set("Content-Type", "text/html; charset-utf-8")
	w.WriteHeader(http.StatusOK)
	buffer.WriteTo(w)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request : %s %s", r.Method, r.URL.Path)

	renderTemplate(w, "index", nil)
}