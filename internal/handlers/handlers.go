package handlers

import "html/template"

type Handler struct {
	Templates map[string]*template.Template
}
