package handlers

import "net/http"

type AuthHandlers struct{}

func NewAuthHandlers() *AuthHandlers {
	return &AuthHandlers{}
}

func (h *AuthHandlers) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: implement login
	w.Write([]byte("login"))
}

func (h *AuthHandlers) Register(w http.ResponseWriter, r *http.Request) {
	// TODO: implement register
	w.Write([]byte("register"))
}

func (h *AuthHandlers) Refresh(w http.ResponseWriter, r *http.Request) {
	// TODO: implement register
	w.Write([]byte("refresh"))
}
