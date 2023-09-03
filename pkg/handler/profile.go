package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/ilgiz-ayupov/auth-app"
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

func (h *Handler) addPhoneNumber(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		sendErrorJSON(w, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	var phoneNumber auth.PhoneNumber

	if err := json.NewDecoder(req.Body).Decode(&phoneNumber); err != nil {
		sendErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	cookie, _ := req.Cookie("SESSTOKEN")
	userClaims, err := h.services.ParseJWTToken(cookie.Value)
	if err != nil {
		sendErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	phoneNumber.UserId = userClaims.UserId

	validate := validator.New()
	if err := validate.Struct(phoneNumber); err != nil {
		sendErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	if _, err := h.services.AddPhoneNumber(phoneNumber); err != nil {
		sendErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
}
