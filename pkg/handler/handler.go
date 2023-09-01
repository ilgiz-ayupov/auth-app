package handler

import (
	"net/http"

	"github.com/ilgiz-ayupov/auth-app/pkg/service"
	"github.com/justinas/alice"
)

type Handler struct {
	services *service.Service
}

func InitHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/user/register", h.userRegister)
	mux.HandleFunc("/user/auth", h.userAuth)

	return alice.New(h.LoggerMiddleware).Then(mux)
}
