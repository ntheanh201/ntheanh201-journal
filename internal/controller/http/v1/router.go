package v1

import (
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"ntheanh201-journal/internal/usecase"
)

func NewRouter(handler *gin.Engine, l logger.Interface, p usecase.Page) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/health", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})

	h := handler.Group("/v1")
	{
		NewPageRoutes(h, p, l)
	}
}
