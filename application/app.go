package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	app := &App{
		rdb: redis.NewClient(&redis.Options{}),
	}
	app.loadRoutes()
	return app
}
func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}
	// pinging to redis
	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect redis: %w", err)
	}
	fmt.Println("Starting Server on https://localhost:3000/")
	ch := make(chan error, 1)
	go func(chError chan error) {
		err = server.ListenAndServe()
		if err != nil {
			chError <- fmt.Errorf("failed to start server: %w", err)
		}
		close(chError)
	}(ch)
	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}

}
