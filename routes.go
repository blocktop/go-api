package api

import (
	"net/http"

	"github.com/blocktop/go-api/handlers"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router gin.IRouter) {
	apiRoute := router.Group("/api")
	apiRoute.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	apiRoute.GET("/metrics/kernel/", handlers.MetricsKernelHandler)

	apiRoute.GET("/metrics/consensus/", handlers.MetricsConsensusHandler)
	apiRoute.GET("/metrics/consensus/tree", handlers.MetricsConsensusTreeHandler)
}