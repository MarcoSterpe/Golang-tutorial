package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
	config Config
}

func New(config Config) *App {
	app := &App{
		rdb: redis.NewClient(&redis.Options{
			Addr: config.RedisAddress,
		}),
		config: config,
	}

	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.ServerPort),
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println(`
Starting server...

====================================================================
________            .___                       _____ __________.___ 
\_____  \_______  __| _/___________  ______   /  _  \\______   \   |
 /   |   \_  __ \/ __ |/ __ \_  __ \/  ___/  /  /_\  \|     ___/   |
/    |    \  | \/ /_/ \  ___/|  | \/\___ \  /    |    \    |   |   |
\_______  /__|  \____ |\___  >__|  /____  > \____|__  /____|   |___|
        \/           \/    \/           \/          \/              
====================================================================
	`)

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		_, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		return server.Shutdown(ctx)
	}
}
