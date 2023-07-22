package service

import (
	"github.com/davidbenavidez/_git/go-gin-microservice-boilerplate.git/internal/clients/db"
)

type Service interface {
	DoSomething
}

type service struct {
	dbClient db.SomeDBClient
}

func New(
	dbClient db.SomeDBClient,
) Service {
	return &service{
		dbClient,
	}
}
