package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/ilgiz-ayupov/auth-app"
)

func (h *Handler) userRegister(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		sendResponseTemplate(w, "template/register-form.html")
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Преобразование возраста в целое число
		age, err := strconv.Atoi(req.PostFormValue("age"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Валидация данных
		user := auth.User{
			Login:    req.PostFormValue("login"),
			Password: req.PostFormValue("password"),
			Name:     req.PostFormValue("name"),
			Age:      age,
		}

		validate := validator.New()
		if err := validate.Struct(user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Создание пользователя
		userId, createdErr := h.services.CreateUser(user)
		if createdErr != nil {
			http.Error(
				w,
				fmt.Sprintf("error creation user: %s", createdErr.Error()),
				http.StatusInternalServerError,
			)
			return
		}

		// Успешная регистрация
		sendResponseHTTP(
			w,
			fmt.Sprintf("Успешная регистрация! Ваш ID в базе данных - %d.", userId),
			http.StatusOK,
		)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) userAuth(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		sendErrorJSON(w, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Чтение данных
	var authFields auth.UserAuthFields

	err := json.NewDecoder(req.Body).Decode(&authFields)
	if err != nil {
		sendErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Создание JWT токена
	token, err := h.services.GenerateJWTToken(authFields.Login, authFields.Password)
	if err != nil {
		sendErrorJSON(w, err, http.StatusUnauthorized)
		return
	}

	// успешная авторизация
	http.SetCookie(w, &http.Cookie{
		Name:     "SESSTOKEN",
		Value:    token,
		Expires:  time.Now().Add(1 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	})

	response := map[string]interface{}{
		"token": token,
	}
	sendResponseJSON(w, response, http.StatusOK)
}
