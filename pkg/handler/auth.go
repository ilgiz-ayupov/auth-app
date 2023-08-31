package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) userRegister(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Regisnter form")
}

func (h *Handler) userAuth(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Auth form")
}
