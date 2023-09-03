package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/ilgiz-ayupov/auth-app"
)

func (h *Handler) getUser(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := mux.Vars(req)["name"]

	user, err := h.services.GetUser(name)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendResponse(w, user, http.StatusOK)
}

func (h *Handler) phoneNumberHandlers(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.searchPhoneNumber(w, req)
	case "POST":
		h.addPhoneNumber(w, req)
	case "PUT":
		h.updatePhoneNumber(w, req)
	}
}

func (h *Handler) searchPhoneNumber(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := req.ParseForm(); err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	phone := req.FormValue("q")
	if len(phone) == 0 {
		sendError(w, "'q' parameter is missing", http.StatusBadRequest)
		return
	}

	phoneNumber, err := h.services.SearchPhoneNumbers(phone)
	if err != nil {
		sendError(w, err.Error(), http.StatusNotFound)
		return
	}

	sendResponse(w, phoneNumber, http.StatusOK)
}

func (h *Handler) addPhoneNumber(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var phoneNum auth.PhoneNumber

	if err := json.NewDecoder(req.Body).Decode(&phoneNum); err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	cookie, _ := req.Cookie("SESSTOKEN")
	userClaims, err := h.services.ParseJWTToken(cookie.Value)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	phoneNum.UserId = userClaims.UserId

	validate := validator.New()
	if err := validate.Struct(phoneNum); err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := h.services.AddPhoneNumber(phoneNum); err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendMessage(w, "Номер телефона был успешно добавлен", http.StatusCreated)
}

func (h *Handler) updatePhoneNumber(w http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var updating auth.UpdatingPhoneNumber

	if err := json.NewDecoder(req.Body).Decode(&updating); err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	cookie, _ := req.Cookie("SESSTOKEN")
	userClaims, err := h.services.ParseJWTToken(cookie.Value)
	if err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updating.UserId = userClaims.UserId

	validate := validator.New()
	if err := validate.Struct(updating); err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.services.UpdatePhoneNumber(updating); err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendMessage(w, "Номер телефона был успешно обновлён", http.StatusOK)
}

func (h *Handler) deletePhoneNumber(w http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	phoneId := mux.Vars(req)["phoneId"]
	phoneNumberId, err := strconv.Atoi(phoneId)
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.services.DeletePhoneNumber(phoneNumberId); err != nil {
		sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendMessage(w, "Номер телефона был успешно удалён", http.StatusOK)
}
