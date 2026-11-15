package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var _temp = template.Must(template.ParseGlob("./templates/*.html"))

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/{$}", homeHandler)

	log.Printf("Serveur: http://localhost:8080")

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatalf("Erreur côté serveur %s", err)
		return
	}

}

func register

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hahaha")
}