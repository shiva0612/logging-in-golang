package main

import (
	"net/http"

	"shiva/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", ping)
	logger.Info("application started")
	router.Run(":8080")

}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
