package handler

import (
	"bytes"
	"log/slog"
	"net/http"

	"lab-quade.de/serverManager/internal/config"
)

func (h *Handler) indexHandler(w http.ResponseWriter, r *http.Request) {
	/* data, err := h.Data.GetAllData() */
	/* if err != nil { */
	/* 	h.ErrorLogger.Error(err.Error()) */
	/* 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) */
	/* } */

	tmpl, ok := h.TemplateCache["index"]
	if !ok {
		h.ErrorLogger.Info(
			"Template does not exist:",
			slog.String("Template", "index"),
		)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	templateData := struct {
		Config *config.Config
	}{
		h.Config,
	}
	buffer := new(bytes.Buffer)
	err := tmpl.ExecuteTemplate(buffer, "base", templateData)
	if err != nil {
		h.ErrorLogger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	buffer.WriteTo(w)
}
