package main

import (
	"net/http"

	"shiva/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	router := gin.Default()

	router.GET("/", ping)
	logger.Info("application started", zap.String("additional_key_runtime", "value"))
	router.Run(":8080")

}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
