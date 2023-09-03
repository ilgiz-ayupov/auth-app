package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilgiz-ayupov/auth-app/pkg/service"
	"github.com/justinas/alice"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/user/register", h.userRegister)
	router.HandleFunc("/user/auth", h.userAuth)

	router.HandleFunc("/user/phone", h.addPhoneNumber).Handler(
		alice.New(h.AuthMiddleware).ThenFunc(h.addPhoneNumber),
	)
	router.HandleFunc("/user/{name}", h.getUser).Handler(
		alice.New(h.AuthMiddleware).ThenFunc(h.getUser),
	)

	return alice.New(h.LoggerMiddleware).Then(router)
}
