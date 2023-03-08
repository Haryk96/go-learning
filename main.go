package main
import (
	"github.com/gin-gonic/gin"
	"jwt-server/router"
)

func main() {
	r := gin.Default()
	router.SetupRouter(r)
	r.Run(":8080")
}
