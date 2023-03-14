package router

import (
	"jwt-server/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// Ping test
	r.GET("/ping", handler.PongHandler)

	r.POST("/jwt", handler.JWTHandler)
}
