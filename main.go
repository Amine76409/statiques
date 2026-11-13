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
