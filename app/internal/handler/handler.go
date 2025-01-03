package handler

import (
	"github.com/arakhimiy/edu-connect/internal/config"
	"github.com/arakhimiy/edu-connect/internal/service"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
	"log/slog"
)

type Handler struct {
	logger  *slog.Logger
	service service.I
	cfg     *config.Config
}

func (h *Handler) Register(router *router.Router[*core.RequestEvent]) {
	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.GET("/test-handler", h.TestHandler)
		}
	}
}

func NewHandler(logger *slog.Logger, service service.I, cfg *config.Config) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
		cfg:     cfg,
	}
}
