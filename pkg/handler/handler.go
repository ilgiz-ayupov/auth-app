package handler

import "net/http"

type Handler struct {
}

func (h *Handler) InitRoutes() {
	http.HandleFunc("/user/register", h.userRegister)
	http.HandleFunc("/user/auth", h.userAuth)
}
