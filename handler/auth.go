package handler

import "net/http"

type Auth struct{}

func (h *Auth) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Works"))
}
func (h *Auth) Signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Signup Works"))
}
func (h *Auth) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUser Works"))
}
