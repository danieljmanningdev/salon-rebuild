package main

import (
	"html/template"
	"log"
	"net/http"

	"salon-rebuild/internal/handlers"
	"salon-rebuild/internal/routes"
)

func main() {
	homeTemplate := template.Must(template.ParseFiles(
		"templates/layouts/base.html",
		"templates/components/header.html",
		"templates/components/footer.html",
		"templates/pages/home.html",
	))

	h := &handlers.Handler{
		Templates: map[string]*template.Template{
			"home": homeTemplate,
		},
	}

	mux := routes.Routes(h)

	log.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
