package main

import (
	"net/http"

	"github.com/davidbenavidez/chi-gorm/internal/config"
	"github.com/davidbenavidez/chi-gorm/internal/log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%s\n", err)
	}
}

func run() error {
	s, port, err := config.SetupServer()

	if err != nil {
		return err
	}

	log.Infof("Listening on port %s", port)
	err = http.ListenAndServe(port, s)

	return err
}
