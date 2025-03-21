package main

import (
	"fmt"
	"html/template"

	assets "lab-quade.de/serverManager"
)

func buildTemplateCache() (map[string]*template.Template, error) {
	templates := assets.Templates()
	cache := map[string]*template.Template{}

	/* PAGES */

	// index
	files := []string{
		"ui/html/base.html",
		"ui/html/pages/index.html",
	}
	tmpl, err := template.ParseFS(templates, files...)
	if err != nil {
		return nil, fmt.Errorf("Kann Index-Template nicht cachen: %w", err)
	}
	cache["index"] = tmpl

	return cache, nil
}
