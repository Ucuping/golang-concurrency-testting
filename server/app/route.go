package app

import (
	"go-concurrency-testting/server/controller"
	"go-concurrency-testting/server/middleware"
)

func route() {
	router.Use(middleware.CORSMiddleware()) // to enable api request client and server

	router.POST("/username", controller.Username)
}
