package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloWorld struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func SetupRouter(r *gin.Engine) {
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/jwt", func(c *gin.Context) {
		helloworld := HelloWorld{
			Name:  "Hello",
			Email: "World",
		}
		c.JSON(http.StatusCreated, helloworld)
	})
}
