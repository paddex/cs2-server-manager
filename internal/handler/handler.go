package handler

import (
	"html/template"
	"log"
	"log/slog"

	"lab-quade.de/serverManager/internal/config"
)

type Handler struct {
	ErrorLogger   *slog.Logger
	AccessLogger  *log.Logger
	TemplateCache map[string]*template.Template
	Config        *config.Config
}
