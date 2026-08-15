package handlers

import (
	"net/http"
	"time"
)

type PageData struct {
	Title       string
	Description string
	Year        int
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:       "Salon",
		Description: "Premium hair and beauty services.",
		Year:        time.Now().Year(),
	}

	tmpl, ok := h.Templates["home"]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
