package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Ethea2/popio-server/database"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":4000",
		Handler: a.router,
	}

	err := db.ConnectDatabase()
	if err != nil {
		fmt.Errorf("something went wrong with app: %w", err)
	}

	fmt.Println("Server running on localhost:4000")
	err = server.ListenAndServe()
	if err != nil {
		fmt.Errorf("something went wrong with app: %w", err)
	}
	return nil
}
