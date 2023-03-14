package handler

import (
	"io"
	"net/http"

	"jwt-server/logger"
	"jwt-server/models"

	"github.com/gin-gonic/gin"
)

func PongHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func JWTHandler(c *gin.Context) {
	cl := logger.Logger
	name := c.Query("name")
	age := c.Query("age")

	var requestBody models.UserCredentials
	if err := c.BindJSON(&requestBody); err != nil {
		if err == io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is missing one or more required field"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cl.Printf("%s, %s\n", name, age)
	cl.Printf("%+v\n", requestBody)
	c.JSON(http.StatusCreated, requestBody)
}
