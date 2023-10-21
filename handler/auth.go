package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ThisJohan/ChitChat/repository"
	"github.com/ThisJohan/ChitChat/utils/jwt"
	"github.com/ThisJohan/ChitChat/utils/password"
)

type Auth struct {
	Repo repository.UserRepo
}

type TokenResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *Auth) Login(w http.ResponseWriter, r *http.Request) {

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &repository.User{}

	if err := h.Repo.FindUserByEmail(user, body.Email).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	if err := password.Verify(user.Password, body.Password); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Email or Password is wrong"))
		return
	}

	token := jwt.Generate(jwt.TokenPayload{ID: user.ID})

	res, _ := json.Marshal(&TokenResponse{Token: token})

	w.Write([]byte(res))

}
func (h *Auth) Signup(w http.ResponseWriter, r *http.Request) {

	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := &repository.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: password.Generate(body.Password),
	}

	if err := h.Repo.CreateUser(user).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token := jwt.Generate(jwt.TokenPayload{ID: user.ID})

	res, _ := json.Marshal(&TokenResponse{Token: token})

	w.Write(res)
	w.WriteHeader(http.StatusCreated)
}
func (h *Auth) GetUser(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header.Get("Authorization"), " ")[1]

	payload, err := jwt.Verify(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user := &repository.User{}

	if err := h.Repo.FindUser(user, "id = ?", payload.ID).Error; err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	res, _ := json.Marshal(&UserResponse{
		Name:  user.Name,
		Email: user.Email,
	})

	w.Write(res)
}
