package handler

import "net/http"

func (h *Handler) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if h.Config.Logging.EnableAccessLog {
			h.AccessLogger.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		}
		next.ServeHTTP(w, r)
	})
}
