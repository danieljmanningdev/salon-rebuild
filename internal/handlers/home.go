package handlers

import "net/http"

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	h.renderPage(
		w,
		"home",
		PageData{
			Title:       "Salon",
			Description: "A fictional contemporary salon concept focused on modern hair and beauty services.",
		},
	)
}
