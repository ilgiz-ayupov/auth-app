package handler

import (
	"net/http"

	"github.com/gorilla/mux"
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
	router := mux.NewRouter()

	router.HandleFunc("/user/register", h.userRegister)
	router.HandleFunc("/user/auth", h.userAuth)

	router.HandleFunc("/user/{name}", h.getUser)

	return alice.New(h.middlewares.LoggerMiddleware).Then(router)
}
