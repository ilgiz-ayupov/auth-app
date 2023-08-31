package handler

import (
	"net/http"

	"github.com/ilgiz-ayupov/auth-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func InitHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() {
	http.HandleFunc("/user/register", h.userRegister)
	http.HandleFunc("/user/auth", h.userAuth)
}
