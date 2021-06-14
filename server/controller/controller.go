package controller

import (
	"errors"
	"go-concurrency-testting/server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Username(c *gin.Context) {
	var urls []string
	if err := c.ShouldBindJSON(&urls); err != nil {
		c.JSON(http.StatusUnprocessableEntity, errors.New("Invalid JSON Body"))
		return
	}
	matchedurls := service.UsernameService.UsernameCheck(urls)

	c.JSON(http.StatusOK, matchedurls)
}
