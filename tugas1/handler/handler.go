package handler

import (
	"github.com/gin-gonic/gin"
)

// handler holds the structure for Handler
type handler struct {}

// Handler holds the contract for Handler
type Handler interface {
	// Ping should handle healthcheck in top level routing
	Ping(*gin.Context)
}

// New is a function for creating handler
func New() Handler {
	return &handler{}
}


// Ping if a function for handling healthcheck in top level routing
func (h *handler) Ping(ctx *gin.Context) {
	select {
	case <-ctx.Request.Context().Done():
		ctx.JSON(408, metaContextTimeout.Message)
		return
	default:
	}

	ctx.JSON(200, gin.H{
		"pingReturn": 1,
	})

	return
}
