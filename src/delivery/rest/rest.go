package rest

import (
	"github.com/dalikewara/ayapingping-go-crud/src/service"
	"github.com/gin-gonic/gin"
)

type REST interface {
	// Serve serves REST application.
	Serve(port int) error
	// RegisterRoutes registers REST routes.
	RegisterRoutes()
}

type Config struct {
	Env                string
	IsProduction       bool
	DefaultOkCode      string
	DefaultOkMessage   string
	DefaultOkStatus    int
	DefaultNotOkStatus int
}

type Handler struct {
	Gin REST
}

type NewParam struct {
	Service   *service.Service
	Config    *Config
	GinClient *gin.Engine
}

// New generates new REST handler.
func New(param NewParam) *Handler {
	return &Handler{
		Gin: NewGin(NewGinParam{
			Service: param.Service,
			Client:  param.GinClient,
			Config:  param.Config,
		}),
	}
}
