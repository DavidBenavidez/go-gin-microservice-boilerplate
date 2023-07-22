package db

import "encoding/json"

type (
	SomeDBClient interface {
		Create() (json.RawMessage, error)
		Read() (json.RawMessage, error)
		Update() (json.RawMessage, error)
		Delete() (json.RawMessage, error)
	}

	client struct {
		host     string
		port     string
		username string
		password string
	}

	createResult struct {
		ID string `json:"id"`
	}
)

func (c *client) Create() (json.RawMessage, error) {
	dbCallReturns := createResult{
		ID: "1",
	}
	return json.Marshal(dbCallReturns)
}
func (c *client) Read() (json.RawMessage, error)   { return nil, nil }
func (c *client) Update() (json.RawMessage, error) { return nil, nil }
func (c *client) Delete() (json.RawMessage, error) { return nil, nil }

func New(
	host,
	port,
	username,
	password string,
) (SomeDBClient, error) {
	// connect to db
	return &client{}, nil
}
