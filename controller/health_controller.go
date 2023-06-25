package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController interface {
	GetHealth(ctx *gin.Context)
}

type healthController struct {
}

func NewHealthController() HealthController {
	return &healthController{}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Tags Health-Controller
// @Accept */*
// @Produce json
// @Success 200
// @Router /health [get]
func (c *healthController) GetHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Up")
}
