package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/ilgiz-ayupov/auth-app"
)

func (h *Handler) userRegister(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		template, err := template.ParseFiles("template/register-form.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := template.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case "POST":
		if err := req.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		age, convErr := strconv.Atoi(req.Form.Get("age"))
		if convErr != nil {
			http.Error(w, convErr.Error(), http.StatusInternalServerError)
			return
		}

		user := auth.User{
			Login:    req.Form.Get("login"),
			Password: req.Form.Get("password"),
			Name:     req.Form.Get("name"),
			Age:      age,
		}

		_, err := h.services.CreateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) userAuth(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Auth form")
}
