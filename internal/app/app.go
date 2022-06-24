package app

import (
	"fmt"
	"github.com/evrone/go-clean-template/pkg/httpserver"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"ntheanh201-journal/config"
	v1 "ntheanh201-journal/internal/controller/http/v1"
	"ntheanh201-journal/internal/usecase"
	"ntheanh201-journal/internal/usecase/webapi"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	pageUseCase := usecase.New(webapi.New())

	handler := gin.New()
	v1.NewRouter(handler, l, pageUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.Http.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
