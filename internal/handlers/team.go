package handlers

import "net/http"

func (h *Handler) Team(w http.ResponseWriter, r *http.Request) {
	h.renderPage(
		w,
		"team",
		PageData{
			Title:       "Team | Salon",
			Description: "Meet the fictional creative team behind this contemporary salon portfolio concept.",
		},
	)
}
