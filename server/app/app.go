package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	route()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	router.Run(":" + port)
}
