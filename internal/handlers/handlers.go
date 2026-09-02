package handlers

import (
	"net/http"
	"time"
)

type Handler struct {
	Templates map[string]*template.Template
}

type PageData struct {
	Title       string
	Description string
	Year        int
}

func (h *Handler) renderPage(
	w http.ResponseWriter,
	key string,
	data PageData,
) {
	data.Year = time.Now().Year()

	tmpl, ok := h.Templates[key]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
