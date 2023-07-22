package rest

import (
	"net/http"

	s "github.com/davidbenavidez/_git/go-gin-microservice-boilerplate.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
)

type rest struct {
	service s.Service
}

func New(
	service s.Service,
	f *fizz.Fizz,
) {
	r := &rest{
		service,
	}

	r.routes(f)
}

func (r *rest) routes(f *fizz.Fizz) {
	infos := &openapi.Info{
		Title:       "Some Service",
		Description: "This service does things",
		Version:     "1.0.0",
	}

	// health endpoint
	f.Engine().GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})

	root := f.Group("/something/v1", "Root endpoint", "Root endpoint")
	root.GET("/swagger/doc.json", nil, f.OpenAPI(infos, "json"), func(context *gin.Context) {})

	root.POST("/:id", []fizz.OperationOption{
		fizz.Description("Create Something"),
		fizz.ID("CreateSomething"),
	}, tonic.Handler(r.something, http.StatusCreated))
}
