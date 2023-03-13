package handler

import (
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

	// Read request body return 400
	// if it does not comply to my model
	var requestBody models.UserCredentials
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cl.Println("Test cl")
	cl.Printf("%s, %s\n", name, age)
	cl.Printf("%+v\n", requestBody)
	c.JSON(http.StatusCreated, requestBody)
}
