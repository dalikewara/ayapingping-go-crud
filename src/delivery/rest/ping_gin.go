package rest

import (
	"github.com/gin-gonic/gin"
)

// Ping handles gin-gonic REST route to ping the server.
func (h *ginHandler) Ping(g *gin.Context) {
	h.jsonOk(g, struct {
		Ping string
	}{
		Ping: "pong",
	})
}
