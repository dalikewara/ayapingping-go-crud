package rest

import (
	"context"
	"fmt"
	"github.com/dalikewara/ayapingping-go-crud/src/entity"
	"github.com/dalikewara/ayapingping-go-crud/src/service"
	"github.com/dalikewara/rflgo"
	"github.com/gin-gonic/gin"
	"log"
)

type ginHandler struct {
	service *service.Service
	client  *gin.Engine
	config  *Config
}

type NewGinParam struct {
	Service *service.Service
	Client  *gin.Engine
	Config  *Config
}

// NewGin generates new gin-gonic REST handler.
func NewGin(param NewGinParam) REST {
	return &ginHandler{
		service: param.Service,
		client:  param.Client,
		config:  param.Config,
	}
}

// RegisterRoutes registers gin-gonic REST routes.
func (h *ginHandler) RegisterRoutes() {
	h.client.POST(RouteUserRegister, h.UserRegister)
	h.client.POST(RouteUserLogin, h.UserLogin)
}

// Serve serves gin-gonic REST application.
func (h *ginHandler) Serve(port int) error {
	log.Println(fmt.Sprintf("REST start on: env=%s, addr=%v, engine=%s", h.config.Env, port, "gin-gonic"))
	if h.config.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	}
	return h.client.Run(fmt.Sprintf(":%v", port))
}

// getRequestContext gets request context.
func (h *ginHandler) getRequestContext(g *gin.Context) context.Context {
	return g.Request.Context()
}

// jsonOk writes success JSON payload into the response body.
func (h *ginHandler) jsonOk(g *gin.Context, data interface{}) {
	g.JSON(Ok(h.config, data))
}

// jsonNotOk writes error JSON payload into the response body.
func (h *ginHandler) jsonNotOk(g *gin.Context, httpStatus int, code, message string) {
	g.JSON(NotOk(h.config, httpStatus, code, message))
}

// jsonCompose composes json response data and writes JSON payload into the response body.
func (h *ginHandler) jsonCompose(g *gin.Context, dest, source interface{}) {
	if err := rflgo.Compose(dest, source); err != nil {
		h.jsonNotOk(g, ErrComposeResponseData.GetStatus(), ErrComposeResponseData.GetCode(), ErrComposeResponseData.GetMessage())
		return
	}
	h.jsonOk(g, dest)
}

// parseJSONRequest parses JSON request into `dest`.
func (h *ginHandler) parseJSONRequest(g *gin.Context, dest any) entity.StdError {
	if err := g.ShouldBindJSON(&dest); err != nil {
		return ErrParseJSONRequest
	}
	return nil
}
