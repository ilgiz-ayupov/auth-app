package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilgiz-ayupov/auth-app/pkg/service"
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
	router.Use(h.LoggerMiddleware)

	router.HandleFunc("/user/register", h.userRegister).Methods("POST")
	router.HandleFunc("/user/auth", h.userAuth).Methods("POST")

	authenticatedRouter := router.PathPrefix("/").Subrouter()
	authenticatedRouter.Use(h.AuthMiddleware)

	authenticatedRouter.HandleFunc("/user/phone", h.phoneNumberHandlers).Methods("GET", "POST", "PUT")
	authenticatedRouter.HandleFunc("/user/phone/{phoneId}", h.deletePhoneNumber)
	authenticatedRouter.HandleFunc("/user/{name}", h.getUser).Methods("GET")

	return router
}
