package handler

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

func (h *Handler) getUser(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		sendErrorJSON(w, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	name := mux.Vars(req)["name"]

	user, err := h.services.GetUser(name)
	if err != nil {
		sendErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		sendErrorJSON(w, errors.New("user not found"), http.StatusBadRequest)
		return
	}

	sendResponseJSON(w, user, http.StatusOK)
}
