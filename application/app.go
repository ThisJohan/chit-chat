package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type App struct {
	router http.Handler
	db     *gorm.DB
}

func New() *App {
	app := &App{}

	app.connectDB()
	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}

	var err error

	fmt.Println("Server Starting")

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to listen to server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
