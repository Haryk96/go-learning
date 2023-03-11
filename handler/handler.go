package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PongHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func JWTHandler(c *gin.Context) {
	type HelloWorld struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	helloworld := HelloWorld{
		Name:  "Hello",
		Email: "World",
	}
	c.JSON(http.StatusCreated, helloworld)
}
