package handlers

import "net/http"

func (h *Handler) Gallery(w http.ResponseWriter, r *http.Request) {
	h.renderPage(
		w,
		"gallery",
		PageData{
			Title:       "Gallery | Salon",
			Description: "Explore a fictional collection of salon looks and beauty treatments created for this portfolio concept.",
		},
	)
}
