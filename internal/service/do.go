package service

import "encoding/json"

type DoSomething interface {
	DoSomething() (json.RawMessage, error)
}

func (s *service) DoSomething() (json.RawMessage, error) {
	// Do something here
	return s.dbClient.Create()
}
