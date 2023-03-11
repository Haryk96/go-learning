package handler

import (
	"net/http"

	"jwt-server/models"

	"github.com/gin-gonic/gin"
)

func PongHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func JWTHandler(c *gin.Context) {

	helloworld := models.HelloWorld{
		Name:  "Hello",
		Email: "World",
	}
	c.JSON(http.StatusCreated, helloworld)
}
