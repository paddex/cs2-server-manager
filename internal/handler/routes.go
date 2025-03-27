package handler

import (
	"net/http"

	assets "lab-quade.de/serverManager"
)

func (h *Handler) Routes() http.Handler {
	staticAssets := assets.Assets()
	fileServer := http.FileServer(http.FS(staticAssets))

	router := http.NewServeMux()
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("GET /{$}", h.indexHandler)
	router.HandleFunc("POST /rcon", h.rconHandler)
	router.HandleFunc("POST /serverinfo", h.serverInfoHandler)

	return h.logRequest(router)
}
