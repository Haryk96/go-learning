package main

import (
	"jwt-server/argparser"
	"jwt-server/router"

	"github.com/gin-gonic/gin"
)

func main() {
	args := argparser.GetArgs()
	r := gin.Default()
	router.SetupRouter(r)
	r.Run(":" + args.Port)
}
