package application

import (
	"net/http"

	"github.com/ThisJohan/ChitChat/handler"
	"github.com/ThisJohan/ChitChat/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/auth", a.loadAuthRoutes)

	a.router = router
}

func (a *App) loadAuthRoutes(r chi.Router) {
	authHandler := &handler.Auth{
		Repo: repository.UserRepo{DB: a.db},
	}

	r.Post("/signup", authHandler.Signup)
	r.Post("/login", authHandler.Login)
	r.Get("/user", authHandler.GetUser)
}
