package main

import (
	"html/template"
	"log"
	"net/http"

	"salon-rebuild/internal/handlers"
	"salon-rebuild/internal/routes"
)

func main() {
	loadPage := func(page string) *template.Template {
		return template.Must(template.ParseFiles(
			"templates/layouts/base.html",
			"templates/components/header.html",
			"templates/components/footer.html",
			"templates/pages/"+page+".html",
		))
	}

	h := &handlers.Handler{
		Templates: map[string]*template.Template{
			"home":     loadPage("home"),
			"services": loadPage("services"),
			"team":     loadPage("team"),
			"gallery":  loadPage("gallery"),
			"contact":  loadPage("contact"),
		},
	}

	mux := routes.Routes(h)

	log.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
