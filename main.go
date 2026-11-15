package main

import (
	"html/template"
	"log"
	"net/http"
	"bytes"
	"strconv"
)

var _temp = template.Must(template.ParseGlob("./templates/*.html"))

type Projects struct {
	Id int
	ProjectName string
	SecretText string
	ShortText string
	Desc string
	Desc2 string
	Link string
}

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/{$}", HomeHandler)
	router.HandleFunc("/projects", ProjectsHandler)
	router.HandleFunc("/projects/{id}", ProjectsByIdHandler)

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

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request : %s %s", r.Method, r.URL.Path)

	renderTemplate(w, "projectIndex", listProjects)
}

var listProjects = []Projects{
	{1, "Projet Rêve", "vraiment déjanté !", "Un rêve qui n'en finira jamais.", "Pour la fin de nos Ymmersions de cette rentrée 2026, on nous a demandé de créer un projet en groupe. Plus précisément, un jeu RPG pûrement textuel où le joueur pourrait parcourir ! <br> Projet Rêve reprend le concept des rêves fiévreux, des rêves où la logique est perturbé ! Mon jeu vous met dans la peau d'un personnage perdu dans un monde où les règles fondamentaux sont brisé.", "Vous venez de vous cracher dans votre vaisseau, sur une planète inconnue. Un chat vous guide, vous apprenez à vous battre, à naviguer dans ce monde. On dirait que vous lisiez un texte sans fin et ni sens, mais, avec votre imagination, vous pouvez vous imaginez à tout.", "https://github.com/vivien-mp3/ProjetRed_CHAUDRON-VIVIEN"},
	{2, "Jeu du pendu", "Est-ce un oiseau ? Est-ce un avion ? Ah.. la blague fini plutôt mal...", "Attachez vos cordes, c'est parti !", "Ce projet simple vient d'une collaboration avec un camarade de classe mais surtout ami, où on a fait un pendu en Golang.", "Ici dans ce projet de classe, vous allez devoir citez des lettres et espérez de trouver le mot, sinon.. Pendu ! Avec un nombre de vies limités, j'espère pour vous que vous êtes prêt !", "https://github.com/Raph-MLT/TP-Jeu-du-pendu-CLI"},
	{3, "Go-Reloaded", "Je vous promets que vos textes seront joli.. enfin je crois-", "Votre texte aura de la prestence !", "Go-reloaded est un projet qui a été donné à toute la classe, le but était de faire un programme un Golang qui permettait de faire la mise en page. <br> Le but étant d'optimiser le fichier texte donné pour en faire un nouveau plus qualitatif !", "Go-Reloaded consiste à faire de la mise en page d'un texte donné par l'utilisateur. le programme va analyser le texte, le corriger et en créer une copie qui est corriger.", "https://github.com/vivien-mp3/Go-Reloaded-CHAUDRON-VIVIEN"},
}


func ProjectsByIdHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request : %s %s", r.Method, r.URL.Path)

	id, errId := strconv.Atoi(r.PathValue("id"))
	if errId != nil {
		http.Error(w, "Erreur données invalides.", http.StatusBadRequest)
		return
	}

	for _, projects := range listProjects {
		log.Println(id, projects.Id)
		if projects.Id == id {
			renderTemplate(w, "projectTemplate", projects)
			return
		}
	}
		http.Error(w, "Erreur projet introuvable.", http.StatusBadRequest)
}