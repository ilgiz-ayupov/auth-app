package handler

import (
	"log"
	"net/http"
)

func (h *Handler) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)
		next.ServeHTTP(w, req)
	})
}
