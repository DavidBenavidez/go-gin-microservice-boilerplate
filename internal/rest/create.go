package rest

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	ID            string `path:"id"`
	Authorization string `header:"authorization"`
}

func (r *rest) something(c *gin.Context, request *CreateRequest) (json.RawMessage, error) {
	// Do something with request
	return r.service.DoSomething()
}
