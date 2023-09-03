package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/ilgiz-ayupov/auth-app"
)

func (h *Handler) userRegister(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := req.ParseForm(); err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	age, err := strconv.Atoi(req.PostFormValue("age"))
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := auth.User{
		Login:    req.PostFormValue("login"),
		Password: req.PostFormValue("password"),
		Name:     req.PostFormValue("name"),
		Age:      age,
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := h.services.CreateUser(user); err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendMessage(w, "Успешная регистрация", http.StatusOK)
}

func (h *Handler) userAuth(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var authFields auth.UserAuthFields

	err := json.NewDecoder(req.Body).Decode(&authFields)
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.services.GenerateJWTToken(authFields.Login, authFields.Password)
	if err != nil {
		sendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "SESSTOKEN",
		Value:    token,
		Expires:  time.Now().Add(1 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	})

	sendMessage(w, "Успешная авторизация", http.StatusOK)
}
