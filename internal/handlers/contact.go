package handlers

import "net/http"

func (h *Handler) Contact(w http.ResponseWriter, r *http.Request) {
	h.renderPage(
		w,
		"contact",
		PageData{
			Title:       "Contact | Salon",
			Description: "Contact the fictional salon concept or start a booking enquiry.",
		},
	)
}
