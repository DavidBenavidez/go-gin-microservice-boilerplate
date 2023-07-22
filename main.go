package main

import (
	"net/http"
	"os"
	"time"

	"github.com/davidbenavidez/_git/go-gin-microservice-boilerplate.git/internal/clients/db"
	"github.com/davidbenavidez/_git/go-gin-microservice-boilerplate.git/internal/rest"
	"github.com/davidbenavidez/_git/go-gin-microservice-boilerplate.git/internal/service"
	"github.com/wI2L/fizz"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		zap.L().Panic("Failed to start server", zap.Error(err))
	}
}

func run() error {

	// Get environment variables
	debug := os.Getenv("DEBUG")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// logger
	setLogger(debug)

	// Make Clients
	dbClient, err := db.New("someHost", "somePort", "someUsername", "somePassword")
	if err != nil {
		zap.L().Error("Error connecting to db", zap.Error(err))
		return err
	}

	// Make service
	svc := service.New(dbClient)

	// Routes
	f := fizz.New()
	rest.New(svc, f)

	srv := &http.Server{
		Addr:              "127.0.0.1:" + port,
		Handler:           f,
		ReadHeaderTimeout: 5 * time.Second,
	}

	zap.L().Info("Listening on port " + port)
	err = srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func setLogger(debug string) {
	var l *zap.Logger

	if len(debug) > 0 {
		l, _ = zap.NewDevelopment()
	} else {
		l, _ = zap.NewProduction()
	}

	defer func() {
		_ = l.Sync()
	}()

	zap.ReplaceGlobals(l)
}
