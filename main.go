package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Projet struct{ ID, Titre, Intro, Depot string }

var projets []Projet

func afficher(w http.ResponseWriter, page string, donnees any) {
	t, _ := template.ParseFiles("templates/" + page)
	t.Execute(w, donnees)
}

func page404(w http.ResponseWriter) {
	w.WriteHeader(404)
	afficher(w, "erreur.html", nil)
}

func accueil(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		page404(w)
		return
	}
	afficher(w, "index.html", nil)
}
