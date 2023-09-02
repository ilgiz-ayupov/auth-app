package handler

import (
	"net/http"

	"github.com/ilgiz-ayupov/auth-app/pkg/service"
	"github.com/justinas/alice"
)

type Handler struct {
	services    *service.Service
	middlewares *Middleware
}

func NewHandler(services *service.Service) *Handler {
	middlewares := NewMiddleware()

	return &Handler{
		services:    services,
		middlewares: middlewares,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/user/register", h.userRegister)
	mux.HandleFunc("/user/auth", h.userAuth)

	return alice.New(h.middlewares.LoggerMiddleware).Then(mux)
}
