package handler

import (
	"fmt"
	"net/http"

	"jwt-server/models"

	"github.com/gin-gonic/gin"
)

func PongHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func JWTHandler(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	// Read request body return 400
	// if it does not comply to my model
	var requestBody models.UserCredentials
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%s, %s", name, age)
	c.JSON(http.StatusCreated, requestBody)
}
