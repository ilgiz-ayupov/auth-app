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

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("SESSTOKEN")
		if err != nil {
			sendError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userTokenClaims, err := h.services.ParseJWTToken(cookie.Value)
		if err := h.services.AuthorizationToken(userTokenClaims); err != nil {
			sendError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, req)
	})
}
