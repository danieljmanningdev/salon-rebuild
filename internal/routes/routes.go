package routes

import (
	"net/http"

	"salon-rebuild/internal/handlers"
)

func Routes(h *handlers.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.Home)
	mux.HandleFunc("/services", h.Services)
	mux.HandleFunc("/team", h.Team)
	mux.HandleFunc("/gallery", h.Gallery)
	mux.HandleFunc("/contact", h.Contact)

	mux.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	return mux
}
