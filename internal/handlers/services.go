package handlers

import "net/http"

func (h *Handler) Services(w http.ResponseWriter, r *http.Request) {
	h.renderPage(
		w,
		"services",
		PageData{
			Title:       "Services | Salon",
			Description: "Explore a fictional menu of contemporary hair and beauty services created for this portfolio concept.",
		},
	)
}
