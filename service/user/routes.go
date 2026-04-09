package user

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /login", h.handleLogin)
	router.HandleFunc("POST /register", h.handleRegister)
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	data := struct {
		Message string `json:"message"`
	}{
		Message: "register successful!",
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic("error")
	}
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	data := struct {
		Message string `json:"message"`
	}{
		Message: "register successful!",
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic("error")
	}
}
