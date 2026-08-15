package routes

import (
	"net/http"

	"salon-rebuild/internal/handlers"
)

func Routes(h *handlers.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.Home)

	mux.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	return mux
}
