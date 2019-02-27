package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gobuffalo/packr/v2"
)

var templatesBox = packr.New("Templates", "../templates")

// HomeHandler / route handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	homeVars := struct {
		Title string
	}{
		Title: "Home",
	}

	templateLayout, err := templatesBox.FindString("layout.html")
	if err != nil {
		log.Fatal(err)
	}
	templateHome, err := templatesBox.FindString("home.html")
	if err != nil {
		log.Fatal(err)
	}
	t := template.New("")
	t.Parse(templateLayout)
	t.Parse(templateHome)
	err = t.ExecuteTemplate(w, "layout", homeVars)
}

// AboutHandler /about route handler
func AboutHandler(w http.ResponseWriter, r *http.Request) {
}
