package handler

import (
	"net/http"

	"gorm.io/gorm"
)

type Auth struct {
	DB *gorm.DB
}

func (h *Auth) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Works"))
}
func (h *Auth) Signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Signup Works"))
}
func (h *Auth) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUser Works"))
}
