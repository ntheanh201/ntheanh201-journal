package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ntheanh201-journal/config"
)

func Run(cfg *config.Config) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
